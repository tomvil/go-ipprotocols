package ipprotocols

// protocolNumbers stores the mapping of protocol numbers to names
var protocolNumbers = map[int]string{
	0:   "HOPOPT",                              // IPv6 Hop-by-Hop Option
	1:   "ICMP",                                // Internet Control Message Protocol
	2:   "IGMP",                                // Internet Group Management Protocol
	3:   "GGP",                                 // Gateway-to-Gateway Protocol
	4:   "IPv4",                                // IPv4 encapsulation
	5:   "ST",                                  // Stream
	6:   "TCP",                                 // Transmission Control Protocol
	7:   "CBT",                                 // Core-Based Trees
	8:   "EGP",                                 // Exterior Gateway Protocol
	9:   "IGP",                                 // Any private interior gateway (used by Cisco for their IGRP)
	10:  "BBN-RCC-MON",                         // BBN RCC Monitoring
	11:  "NVP-II",                              // Network Voice Protocol
	12:  "PUP",                                 // Xerox PUP
	13:  "ARGUS",                               // ARGUS
	14:  "EMCON",                               // EMCON
	15:  "XNET",                                // Cross Net Debugger
	16:  "CHAOS",                               // Chaos
	17:  "UDP",                                 // User Datagram Protocol
	18:  "MUX",                                 // Multiplexing
	19:  "DCN-MEAS",                            // DCN Measurement Subsystems
	20:  "HMP",                                 // Host Monitoring Protocol
	21:  "PRM",                                 // Packet Radio Measurement
	22:  "XNS-IDP",                             // XEROX NS IDP
	23:  "TRUNK-1",                             // Trunk-1
	24:  "TRUNK-2",                             // Trunk-2
	25:  "LEAF-1",                              // Leaf-1
	26:  "LEAF-2",                              // Leaf-2
	27:  "RDP",                                 // Reliable Data Protocol
	28:  "IRTP",                                // Internet Reliable Transaction Protocol
	29:  "ISO-TP4",                             // ISO Transport Protocol Class 4
	30:  "NETBLT",                              // Bulk Data Transfer Protocol
	31:  "MFE-NSP",                             // MFE Network Services Protocol
	32:  "MERIT-INP",                           // MERIT Internodal Protocol
	33:  "DCCP",                                // Datagram Congestion Control Protocol
	34:  "3PC",                                 // Third Party Connect Protocol
	35:  "IDPR",                                // Inter-Domain Policy Routing Protocol
	36:  "XTP",                                 // XTP
	37:  "DDP",                                 // Datagram Delivery Protocol
	38:  "IDPR-CMTP",                           // IDPR Control Message Transport Protocol
	39:  "TP++",                                // TP++ Transport Protocol
	40:  "IL",                                  // IL Transport Protocol
	41:  "IPv6",                                // IPv6 encapsulation
	42:  "SDRP",                                // Source Demand Routing Protocol
	43:  "IPv6-Route",                          // Routing Header for IPv6
	44:  "IPv6-Frag",                           // Fragment Header for IPv6
	45:  "IDRP",                                // Inter-Domain Routing Protocol
	46:  "RSVP",                                // Reservation Protocol
	47:  "GRE",                                 // Generic Routing Encapsulation
	48:  "DSR",                                 // Dynamic Source Routing Protocol
	49:  "BNA",                                 // BNA
	50:  "ESP",                                 // Encapsulating Security Payload
	51:  "AH",                                  // Authentication Header
	52:  "I-NLSP",                              // Integrated Net Layer Security Protocol
	53:  "SWIPE",                               // IP with Encryption
	54:  "NARP",                                // NBMA Address Resolution Protocol
	55:  "MOBILE",                              // IP Mobility
	56:  "TLSP",                                // Transport Layer Security Protocol (using Kryptonet key management)
	57:  "SKIP",                                // Simple Key-Management for Internet Protocol
	58:  "ICMPv6",                              // ICMP for IPv6
	59:  "IPv6-NoNxt",                          // No Next Header for IPv6
	60:  "IPv6-Opts",                           // Destination Options for IPv6
	61:  "Any host internal protocol",          // Any host internal protocol
	62:  "CFTP",                                // CFTP
	63:  "Any local network",                   // Any local network
	64:  "SAT-EXPAK",                           // SATNET and Backroom EXPAK
	65:  "KRYPTOLAN",                           // Kryptolan
	66:  "RVD",                                 // MIT Remote Virtual Disk Protocol
	67:  "IPPC",                                // Internet Pluribus Packet Core
	68:  "Any distributed file system",         // Any distributed file system
	69:  "SAT-MON",                             // SATNET Monitoring
	70:  "VISA",                                // VISA Protocol
	71:  "IPCV",                                // Internet Packet Core Utility
	72:  "CPNX",                                // Computer Protocol Network Executive
	73:  "CPHB",                                // Computer Protocol Heart Beat
	74:  "WSN",                                 // Wang Span Network
	75:  "PVP",                                 // Packet Video Protocol
	76:  "BRSAT-MON",                           // Backroom SATNET Monitoring
	77:  "SUN-ND",                              // SUN ND PROTOCOL-Temporary
	78:  "WB-MON",                              // WIDEBAND Monitoring
	79:  "WB-EXPAK",                            // WIDEBAND EXPAK
	80:  "ISO-IP",                              // International Organization for Standardization Internet Protocol
	81:  "VMTP",                                // Versatile Message Transaction Protocol
	82:  "SECURE-VMTP",                         // Secure Versatile Message Transaction Protocol
	83:  "VINES",                               // VINES
	84:  "TTP",                                 // TTP
	85:  "NSFNET-IGP",                          // NSFNET-IGP
	86:  "DGP",                                 // Dissimilar Gateway Protocol
	87:  "TCF",                                 // TCF
	88:  "EIGRP",                               // EIGRP
	89:  "OSPF",                                // Open Shortest Path First
	90:  "Sprite-RPC",                          // Sprite RPC Protocol
	91:  "LARP",                                // Locus Address Resolution Protocol
	92:  "MTP",                                 // Multicast Transport Protocol
	93:  "AX.25",                               // AX.25
	94:  "IPIP",                                // IP-within-IP Encapsulation Protocol
	95:  "MICP",                                // Mobile Internetworking Control Protocol
	96:  "SCC-SP",                              // Semaphore Communications Sec. Pro
	97:  "ETHERIP",                             // Ethernet-within-IP Encapsulation
	98:  "ENCAP",                               // Encapsulation Header
	99:  "Any private encryption scheme",       // Any private encryption scheme
	100: "GMTP",                                // GMTP
	101: "IFMP",                                // Ipsilon Flow Management Protocol
	102: "PNNI",                                // PNNI over IP
	103: "PIM",                                 // Protocol Independent Multicast
	104: "ARIS",                                // ARIS
	105: "SCPS",                                // SCPS
	106: "QNX",                                 // QNX
	107: "A/N",                                 // Active Networks
	108: "IPComp",                              // IP Payload Compression Protocol
	109: "SNP",                                 // Sitara Networks Protocol
	110: "Compaq-Peer",                         // Compaq Peer Protocol
	111: "IPX-in-IP",                           // IPX in IP
	112: "VRRP",                                // Virtual Router Redundancy Protocol
	113: "PGM",                                 // PGM Reliable Transport Protocol
	114: "Any 0-hop protocol",                  // Any 0-hop protocol
	115: "L2TP",                                // Layer Two Tunneling Protocol
	116: "DDX",                                 // D-II Data Exchange (DDX)
	117: "IATP",                                // Interactive Agent Transfer Protocol
	118: "STP",                                 // Schedule Transfer Protocol
	119: "SRP",                                 // SpectraLink Radio Protocol
	120: "UTI",                                 // UTI
	121: "SMP",                                 // Simple Message Protocol
	122: "SM",                                  // Simple Multicast Protocol
	123: "PTP",                                 // Performance Transparency Protocol
	124: "ISIS over IPv4",                      // ISIS over IPv4
	125: "FIRE",                                // Flexible Intra-AS Routing Environment
	126: "CRTP",                                // Combat Radio Transport Protocol
	127: "CRUDP",                               // Combat Radio User Datagram
	128: "SSCOPMCE",                            // SSCOPMCE
	129: "IPLT",                                // IPLT
	130: "SPS",                                 // Secure Packet Shield
	131: "PIPE",                                // Private IP Encapsulation within IP
	132: "SCTP",                                // Stream Control Transmission Protocol
	133: "FC",                                  // Fibre Channel
	134: "RSVP-E2E-IGNORE",                     // RSVP E2E Ignore
	135: "Mobility Header",                     // Mobility Extension Header for IPv6
	136: "UDPLite",                             // Lightweight User Datagram Protocol
	137: "MPLS-in-IP",                          // MPLS-in-IP
	138: "manet",                               // MANET Protocols
	139: "HIP",                                 // Host Identity Protocol
	140: "Shim6",                               // Site Multihoming by IPv6 Intermediation
	141: "WESP",                                // Wrapped Encapsulating Security Payload
	142: "ROHC",                                // Robust Header Compression
	143: "Ethernet",                            // Ethernet
	144: "AGGFRAG",                             // AGGFRAG encapsulation payload for ESP
	253: "Use for experimentation and testing", // Use for experimentation and testing
	254: "Use for experimentation and testing", // Use for experimentation and testing
	255: "Reserved",                            // Reserved
}

// protocolNames stores the static mapping of protocol names to numbers.
var protocolNames = map[string]int{
	"hopopt":                        0,
	"icmp":                          1,
	"igmp":                          2,
	"ggp":                           3,
	"ipv4":                          4,
	"st":                            5,
	"tcp":                           6,
	"cbt":                           7,
	"egp":                           8,
	"igp":                           9,
	"bbn-rcc-mon":                   10,
	"nvp-ii":                        11,
	"pup":                           12,
	"argus":                         13,
	"emcon":                         14,
	"xnet":                          15,
	"chaos":                         16,
	"udp":                           17,
	"mux":                           18,
	"dcn-meas":                      19,
	"hmp":                           20,
	"prm":                           21,
	"xns-idp":                       22,
	"trunk-1":                       23,
	"trunk-2":                       24,
	"leaf-1":                        25,
	"leaf-2":                        26,
	"rdp":                           27,
	"irtp":                          28,
	"iso-tp4":                       29,
	"netblt":                        30,
	"mfe-nsp":                       31,
	"merit-inp":                     32,
	"dccp":                          33,
	"3pc":                           34,
	"idpr":                          35,
	"xtp":                           36,
	"ddp":                           37,
	"idpr-cmtp":                     38,
	"tp++":                          39,
	"il":                            40,
	"ipv6":                          41,
	"sdrp":                          42,
	"ipv6-route":                    43,
	"ipv6-frag":                     44,
	"idrp":                          45,
	"rsvp":                          46,
	"gre":                           47,
	"dsr":                           48,
	"bna":                           49,
	"esp":                           50,
	"ah":                            51,
	"i-nlsp":                        52,
	"swipe":                         53,
	"narp":                          54,
	"mobile":                        55,
	"tlsp":                          56,
	"skip":                          57,
	"icmpv6":                        58,
	"ipv6-nonxt":                    59,
	"ipv6-opts":                     60,
	"any host internal protocol":    61,
	"cftp":                          62,
	"any local network":             63,
	"sat-expak":                     64,
	"kryptolan":                     65,
	"rvd":                           66,
	"ippc":                          67,
	"any distributed file system":   68,
	"sat-mon":                       69,
	"visa":                          70,
	"ipcv":                          71,
	"cpnx":                          72,
	"cphb":                          73,
	"wsn":                           74,
	"pvp":                           75,
	"brsat-mon":                     76,
	"sun-nd":                        77,
	"wb-mon":                        78,
	"wb-expak":                      79,
	"iso-ip":                        80,
	"vmtp":                          81,
	"secure-vmtp":                   82,
	"vines":                         83,
	"ttp":                           84,
	"nsfnet-igp":                    85,
	"dgp":                           86,
	"tcf":                           87,
	"eigrp":                         88,
	"ospf":                          89,
	"sprite-rpc":                    90,
	"larp":                          91,
	"mtp":                           92,
	"ax.25":                         93,
	"ipip":                          94,
	"micp":                          95,
	"scc-sp":                        96,
	"etherip":                       97,
	"encap":                         98,
	"any private encryption scheme": 99,
	"gmtp":                          100,
	"ifmp":                          101,
	"pnni":                          102,
	"pim":                           103,
	"aris":                          104,
	"scps":                          105,
	"qnx":                           106,
	"a/n":                           107,
	"ipcomp":                        108,
	"snp":                           109,
	"compaq-peer":                   110,
	"ipx-in-ip":                     111,
	"vrrp":                          112,
	"pgm":                           113,
	"any 0-hop protocol":            114,
	"l2tp":                          115,
	"ddx":                           116,
	"iatp":                          117,
	"stp":                           118,
	"srp":                           119,
	"uti":                           120,
	"smp":                           121,
	"sm":                            122,
	"ptp":                           123,
	"isis over ipv4":                124,
	"fire":                          125,
	"crtp":                          126,
	"crudp":                         127,
	"sscopmce":                      128,
	"iplt":                          129,
	"sps":                           130,
	"pipe":                          131,
	"sctp":                          132,
	"fc":                            133,
	"rsvp-e2e-ignore":               134,
	"mobility header":               135,
	"udplite":                       136,
	"mpls-in-ip":                    137,
	"manet":                         138,
	"hip":                           139,
	"shim6":                         140,
	"wesp":                          141,
	"rohc":                          142,
	"ethernet":                      143,
	"aggfrag":                       144,
	"use for experimentation and testing (253)": 253,
	"use for experimentation and testing (254)": 254,
	"reserved": 255,
}
