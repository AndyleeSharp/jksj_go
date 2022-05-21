#### 总结几种 socket 粘包的解包方式：fix length/delimiter based/length field based frame decoder。尝试举例其应用。

#### fix length:
定长数据包，每次接收或发送约定好的固定长度数据，应用比较少。
#### delimiter based：
基于分隔符分包，当读取到特定符号时代表一个包发送或接收结束，应用比如http包的收发。
#### length field based frame decoder
一般利用包头的长度信息让接收方知道包头/包体的长度，每次处理指定长度的数据，从而达到分包的目的。互联网大多数传输协议都使用这种分包方式。

