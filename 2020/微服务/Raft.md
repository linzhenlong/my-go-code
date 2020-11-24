##Raft算法--保证一致性

###一、简介

* Raft提供了一种在计算系统集群中分布式状态机的通用方法，确保集群中的每个节点都统一一系列的状态转换
* 它有许多开源参考实现，具有go,c++,java和scala中的完整规范实现
* 一个raft集群包含若干个服务器节点，通常是5个，这允许整个系统容忍2个节点的失败，每个节点处于三种状态之一
    * follower(跟随者)：所有节点都以follower的状态开始。如果没有收到Leader的消息则会变成candidate状态
    * candidate(候选者)：会向其他节点拉"选票"，如果得到了大部分的票则成为Leader,这个过程叫做Leader选举(Leader Election)
    * Leader(领导者):所有系统的修改都会先经过leader
    
###二、一致性算法

* Raft通过选出一个leader来简化日志副本的管理，例如，日志项（log entry）只允许从leader流向follower
* 基于leader的方法，Raft算法可以分解为三个子问题
    * Leader election(领导选举):原来的leader挂断后，必须选出一个新的leader
    * Log replication(日志复制):leader从客户端收日志，并复制到整个集群中
    * Safety(安全性):如果有任意server将日志回收项放回到状态机中，那么其他的server只会回放相同的日志项
    
###三、Raft动画演示

* 地址：http://thesecretlivesofdata.com/raft/
* 动画主要包含三个部分
    * 第一部分介绍简单版的领导者选举和日志复制的过程
    * 第二部分介绍详细版的领导选举和日志复制的过程
    * 第三部分介绍如果遇到网络分区(脑裂)，raft算法是如何恢复网络一致的
    
 