package sdf

// 定义SDF接口涉及到的对象（参考./sansec/swsds.h文件定义）

// 设备信息
type DeviceInfo struct {
	IssuerName      string	// 设备生产厂商名称
	DeviceName      string	// 设备型号
	DeviceSerial    string	// 设备编号
	DeviceVersion   uint	// 密码设备内部软件版本号
	StandardVersion uint	// 密码设备支持的接口规范版本号
	AsymAlgAbility  [2]uint	// （非对称算法）前四字节表示支持的算法；后四字节表示算法的最大模长
	SymAlgAbility   uint	// （对称算法）所有支持的对称算法
	HashAlgAbility  uint	// 所有支持的杂凑算法
	BufferSize      uint	// 支持的最大文件存储空间
}

type DeviceRunStatus struct {
	Onboot      uint
	Service     uint
	Concurrency uint
	Memtotal    uint
	Memfree     uint
	Cpu         uint
	Reserve1    uint
	Reserve2    uint
}

type RSArefPublicKeyLite struct {
	Bits uint
	M    string
	E    string
}

type RSArefPrivateKeyLite struct {
	Bits  uint
	M     string
	E     string
	D     string
	Prime [2]string
	Pexp  [2]string
	Coef  string
}

type RSArefPublicKey struct {
	Bits uint
	M    string
	E    string
}

type RSArefPrivateKey struct {
	Bits  uint
	M     string
	E     string
	D     string
	Prime [2]string
	Pexp  [2]string
	Coef  string
}

type ECCrefPublicKey struct {
	Bits uint
	X    string
	Y    string
}

type ECCrefPrivateKey struct {
	Bits uint
	K    string
}

type ECCCipher struct {
	X string
	Y string
	M string
	L uint
	C string
}

type ECCSignature struct {
	R string
	S string
}

type SM9refSignMasterPrivateKey struct {
	Bits uint
	S    string
}

type SM9refSignMasterPublicKey struct {
	Bits uint
	Xa   string
	Xb   string
	Ya   string
	Yb   string
}

type SM9refEncMasterPrivateKey struct {
	Bits uint
	S    string
}

type SM9refEncMasterPublicKey struct {
	Bits uint
	X    string
	Y    string
}

type SM9refSignUserPrivateKey struct {
	Bits uint
	X    string
	Y    string
}

type SM9refEncUserPrivateKey struct {
	Bits uint
	Xa   string
	Xb   string
	Ya   string
	Yb   string
}

type SM9Signature struct {
	H string
	X string
	Y string
}

type SM9Cipher struct {
	X string
	Y string
	H string
	L uint
	C string
}

type SM9refKeyPackage struct {
	X string
	Y string
}
