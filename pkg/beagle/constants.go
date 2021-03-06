package beagle

const (
	// BG_DEBUG as defined in beagle/beagle.h:101
	BG_DEBUG = 0
	// BG_HEADER_VERSION as defined in beagle/beagle.h:107
	BG_HEADER_VERSION = 1320
	// BG_PORT_NOT_FREE as defined in beagle/beagle.h:276
)

const BG_PORT_NOT_FREE = 0x8000
const (
	// BG_FEATURE_NONE as defined in beagle/beagle.h:374
	BG_FEATURE_NONE = 0x00000000
	// BG_FEATURE_I2C as defined in beagle/beagle.h:375
	BG_FEATURE_I2C = 0x00000001
	// BG_FEATURE_SPI as defined in beagle/beagle.h:376
	BG_FEATURE_SPI = 0x00000002
	// BG_FEATURE_USB as defined in beagle/beagle.h:377
	BG_FEATURE_USB = 0x00000004
	// BG_FEATURE_MDIO as defined in beagle/beagle.h:378
	BG_FEATURE_MDIO = 0x00000008
	// BG_FEATURE_USB_HS as defined in beagle/beagle.h:379
	BG_FEATURE_USB_HS = 0x00000010
	// BG_FEATURE_USB_SS as defined in beagle/beagle.h:380
	BG_FEATURE_USB_SS = 0x00000020
)
const (
	// BG_TARGET_POWER_OFF as defined in beagle/beagle.h:458
	BG_TARGET_POWER_OFF = 0
	// BG_TARGET_POWER_ON as defined in beagle/beagle.h:459
	BG_TARGET_POWER_ON = 1
	// BG_TARGET_POWER_QUERY as defined in beagle/beagle.h:460
	BG_TARGET_POWER_QUERY = 128
)
const (
	// BG_HOST_IFCE_FULL_SPEED as defined in beagle/beagle.h:467
	BG_HOST_IFCE_FULL_SPEED = 0
	// BG_HOST_IFCE_HIGH_SPEED as defined in beagle/beagle.h:468
	BG_HOST_IFCE_HIGH_SPEED = 1
	// BG_HOST_IFCE_SUPER_SPEED as defined in beagle/beagle.h:469
	BG_HOST_IFCE_SUPER_SPEED = 2
)

const (
	// BG_I2C_MONITOR_DATA as defined in beagle/beagle.h:627
	BG_I2C_MONITOR_DATA = 255
	// BG_I2C_MONITOR_NACK as defined in beagle/beagle.h:628
	BG_I2C_MONITOR_NACK = 256
	// BG_READ_I2C_NO_STOP as defined in beagle/beagle.h:629
	BG_READ_I2C_NO_STOP = 65536
	// BG_USB_PID_OUT as defined in beagle/beagle.h:751
	BG_USB_PID_OUT = 225
	// BG_USB_PID_IN as defined in beagle/beagle.h:752
	BG_USB_PID_IN = 105
	// BG_USB_PID_SOF as defined in beagle/beagle.h:753
	BG_USB_PID_SOF = 165
	// BG_USB_PID_SETUP as defined in beagle/beagle.h:754
	BG_USB_PID_SETUP = 45
	// BG_USB_PID_DATA0 as defined in beagle/beagle.h:755
	BG_USB_PID_DATA0 = 195
	// BG_USB_PID_DATA1 as defined in beagle/beagle.h:756
	BG_USB_PID_DATA1 = 75
	// BG_USB_PID_DATA2 as defined in beagle/beagle.h:757
	BG_USB_PID_DATA2 = 135
	// BG_USB_PID_MDATA as defined in beagle/beagle.h:758
	BG_USB_PID_MDATA = 15
	// BG_USB_PID_ACK as defined in beagle/beagle.h:759
	BG_USB_PID_ACK = 210
	// BG_USB_PID_NAK as defined in beagle/beagle.h:760
	BG_USB_PID_NAK = 90
	// BG_USB_PID_STALL as defined in beagle/beagle.h:761
	BG_USB_PID_STALL = 30
	// BG_USB_PID_NYET as defined in beagle/beagle.h:762
	BG_USB_PID_NYET = 150
	// BG_USB_PID_PRE as defined in beagle/beagle.h:763
	BG_USB_PID_PRE = 60
	// BG_USB_PID_ERR as defined in beagle/beagle.h:764
	BG_USB_PID_ERR = 60
	// BG_USB_PID_SPLIT as defined in beagle/beagle.h:765
	BG_USB_PID_SPLIT = 120
	// BG_USB_PID_PING as defined in beagle/beagle.h:766
	BG_USB_PID_PING = 180
	// BG_USB_PID_EXT as defined in beagle/beagle.h:767
	BG_USB_PID_EXT = 240
	// BG_USB_PID_CORRUPTED as defined in beagle/beagle.h:768
	BG_USB_PID_CORRUPTED = 255
	// BG_READ_USB_ERR_BAD_SIGNALS as defined in beagle/beagle.h:770
	BG_READ_USB_ERR_BAD_SIGNALS = 65536
	// BG_READ_USB_ERR_BAD_PID as defined in beagle/beagle.h:771
	BG_READ_USB_ERR_BAD_PID = 2097152
	// BG_READ_USB_ERR_BAD_CRC as defined in beagle/beagle.h:772
	BG_READ_USB_ERR_BAD_CRC = 4194304
	// BG_READ_USB_ERR_BAD_SYNC as defined in beagle/beagle.h:774
	BG_READ_USB_ERR_BAD_SYNC = 131072
	// BG_READ_USB_ERR_BIT_STUFF as defined in beagle/beagle.h:775
	BG_READ_USB_ERR_BIT_STUFF = 262144
	// BG_READ_USB_ERR_FALSE_EOP as defined in beagle/beagle.h:776
	BG_READ_USB_ERR_FALSE_EOP = 524288
	// BG_READ_USB_ERR_LONG_EOP as defined in beagle/beagle.h:777
	BG_READ_USB_ERR_LONG_EOP = 1048576
	// BG_READ_USB_TRUNCATION_LEN_MASK as defined in beagle/beagle.h:779
	BG_READ_USB_TRUNCATION_LEN_MASK = 255
	// BG_READ_USB_TRUNCATION_MODE as defined in beagle/beagle.h:780
	BG_READ_USB_TRUNCATION_MODE = 536870912
	// BG_READ_USB_END_OF_CAPTURE as defined in beagle/beagle.h:781
	BG_READ_USB_END_OF_CAPTURE = 1073741824
	// BG_READ_USB_ERR_BAD_SLC_CRC_1 as defined in beagle/beagle.h:783
	BG_READ_USB_ERR_BAD_SLC_CRC_1 = 32768
	// BG_READ_USB_ERR_BAD_SLC_CRC_2 as defined in beagle/beagle.h:784
	BG_READ_USB_ERR_BAD_SLC_CRC_2 = 65536
	// BG_READ_USB_ERR_BAD_SHP_CRC_16 as defined in beagle/beagle.h:785
	BG_READ_USB_ERR_BAD_SHP_CRC_16 = 32768
	// BG_READ_USB_ERR_BAD_SHP_CRC_5 as defined in beagle/beagle.h:786
	BG_READ_USB_ERR_BAD_SHP_CRC_5 = 65536
	// BG_READ_USB_ERR_BAD_SDP_CRC as defined in beagle/beagle.h:787
	BG_READ_USB_ERR_BAD_SDP_CRC = 32768
	// BG_READ_USB_EDB_FRAMING as defined in beagle/beagle.h:788
	BG_READ_USB_EDB_FRAMING = 131072
	// BG_READ_USB_ERR_UNK_END_OF_FRAME as defined in beagle/beagle.h:789
	BG_READ_USB_ERR_UNK_END_OF_FRAME = 262144
	// BG_READ_USB_ERR_DATA_LEN_INVALID as defined in beagle/beagle.h:790
	BG_READ_USB_ERR_DATA_LEN_INVALID = 524288
	// BG_READ_USB_PKT_TYPE_LINK as defined in beagle/beagle.h:791
	BG_READ_USB_PKT_TYPE_LINK = 1048576
	// BG_READ_USB_PKT_TYPE_HDR as defined in beagle/beagle.h:792
	BG_READ_USB_PKT_TYPE_HDR = 2097152
	// BG_READ_USB_PKT_TYPE_DP as defined in beagle/beagle.h:793
	BG_READ_USB_PKT_TYPE_DP = 4194304
	// BG_READ_USB_PKT_TYPE_TSEQ as defined in beagle/beagle.h:794
	BG_READ_USB_PKT_TYPE_TSEQ = 8388608
	// BG_READ_USB_PKT_TYPE_TS1 as defined in beagle/beagle.h:795
	BG_READ_USB_PKT_TYPE_TS1 = 16777216
	// BG_READ_USB_PKT_TYPE_TS2 as defined in beagle/beagle.h:796
	BG_READ_USB_PKT_TYPE_TS2 = 33554432
	// BG_READ_USB_ERR_BAD_TS as defined in beagle/beagle.h:797
	BG_READ_USB_ERR_BAD_TS = 67108864
	// BG_READ_USB_ERR_FRAMING as defined in beagle/beagle.h:798
	BG_READ_USB_ERR_FRAMING = 134217728
	// BG_READ_USBPD_BIT_LEN_MASK as defined in beagle/beagle.h:800
	BG_READ_USBPD_BIT_LEN_MASK = 1792
	// BG_READ_USBPD_BIT_LEN_SHIFT as defined in beagle/beagle.h:801
	BG_READ_USBPD_BIT_LEN_SHIFT = 8
	// BG_READ_USBPD_PARSE_ERR_MASK as defined in beagle/beagle.h:802
	BG_READ_USBPD_PARSE_ERR_MASK = 15
	// BG_READ_USBPD_NO_PREAMBLE as defined in beagle/beagle.h:803
	BG_READ_USBPD_NO_PREAMBLE = 1
	// BG_READ_USBPD_NO_SOP as defined in beagle/beagle.h:804
	BG_READ_USBPD_NO_SOP = 2
	// BG_READ_USBPD_NO_HEADER as defined in beagle/beagle.h:805
	BG_READ_USBPD_NO_HEADER = 3
	// BG_READ_USBPD_NO_DATA as defined in beagle/beagle.h:806
	BG_READ_USBPD_NO_DATA = 4
	// BG_READ_USBPD_NO_CRC as defined in beagle/beagle.h:807
	BG_READ_USBPD_NO_CRC = 5
	// BG_READ_USBPD_NO_EOP as defined in beagle/beagle.h:808
	BG_READ_USBPD_NO_EOP = 6
	// BG_READ_USBPD_ERR_BAD_CRC as defined in beagle/beagle.h:809
	BG_READ_USBPD_ERR_BAD_CRC = 16777216
	// BG_READ_USBPD_ERR_UNKNOWN_TYPE as defined in beagle/beagle.h:810
	BG_READ_USBPD_ERR_UNKNOWN_TYPE = 33554432
	// BG_EVENT_USB_HOST_DISCONNECT as defined in beagle/beagle.h:812
	BG_EVENT_USB_HOST_DISCONNECT = 256
	// BG_EVENT_USB_TARGET_DISCONNECT as defined in beagle/beagle.h:813
	BG_EVENT_USB_TARGET_DISCONNECT = 512
	// BG_EVENT_USB_HOST_CONNECT as defined in beagle/beagle.h:814
	BG_EVENT_USB_HOST_CONNECT = 1024
	// BG_EVENT_USB_TARGET_CONNECT as defined in beagle/beagle.h:815
	BG_EVENT_USB_TARGET_CONNECT = 2048
	// BG_EVENT_USB_RESET as defined in beagle/beagle.h:816
	BG_EVENT_USB_RESET = 4096
	// BG_EVENT_USB_DIGITAL_INPUT_MASK as defined in beagle/beagle.h:821
	BG_EVENT_USB_DIGITAL_INPUT_MASK = 15
	// BG_EVENT_USB_CHIRP_J as defined in beagle/beagle.h:822
	BG_EVENT_USB_CHIRP_J = 8192
	// BG_EVENT_USB_CHIRP_K as defined in beagle/beagle.h:823
	BG_EVENT_USB_CHIRP_K = 16384
	// BG_EVENT_USB_SPEED_UNKNOWN as defined in beagle/beagle.h:824
	BG_EVENT_USB_SPEED_UNKNOWN = 32768
	// BG_EVENT_USB_LOW_SPEED as defined in beagle/beagle.h:825
	BG_EVENT_USB_LOW_SPEED = 65536
	// BG_EVENT_USB_FULL_SPEED as defined in beagle/beagle.h:826
	BG_EVENT_USB_FULL_SPEED = 131072
	// BG_EVENT_USB_HIGH_SPEED as defined in beagle/beagle.h:827
	BG_EVENT_USB_HIGH_SPEED = 262144
	// BG_EVENT_USB_LOW_OVER_FULL_SPEED as defined in beagle/beagle.h:828
	BG_EVENT_USB_LOW_OVER_FULL_SPEED = 524288
	// BG_EVENT_USB_SUSPEND as defined in beagle/beagle.h:829
	BG_EVENT_USB_SUSPEND = 1048576
	// BG_EVENT_USB_RESUME as defined in beagle/beagle.h:830
	BG_EVENT_USB_RESUME = 2097152
	// BG_EVENT_USB_KEEP_ALIVE as defined in beagle/beagle.h:831
	BG_EVENT_USB_KEEP_ALIVE = 4194304
	// BG_EVENT_USB_DIGITAL_INPUT as defined in beagle/beagle.h:832
	BG_EVENT_USB_DIGITAL_INPUT = 8388608
	// BG_EVENT_USB_OTG_HNP as defined in beagle/beagle.h:833
	BG_EVENT_USB_OTG_HNP = 16777216
	// BG_EVENT_USB_OTG_SRP_DATA_PULSE as defined in beagle/beagle.h:834
	BG_EVENT_USB_OTG_SRP_DATA_PULSE = 33554432
	// BG_EVENT_USB_OTG_SRP_VBUS_PULSE as defined in beagle/beagle.h:835
	BG_EVENT_USB_OTG_SRP_VBUS_PULSE = 67108864
	// BG_EVENT_USB_SMA_EXTIN_DETECTED as defined in beagle/beagle.h:840
	BG_EVENT_USB_SMA_EXTIN_DETECTED = 134217728
	// BG_EVENT_USB_CHIRP_DETECTED as defined in beagle/beagle.h:841
	BG_EVENT_USB_CHIRP_DETECTED = 128
	// BG_EVENT_USB_VBUS_TRIGGER as defined in beagle/beagle.h:843
	BG_EVENT_USB_VBUS_TRIGGER = 134217728
	// BG_EVENT_USB_COMPLEX_TIMER as defined in beagle/beagle.h:844
	BG_EVENT_USB_COMPLEX_TIMER = 268435456
	// BG_EVENT_USB_CROSS_TRIGGER as defined in beagle/beagle.h:845
	BG_EVENT_USB_CROSS_TRIGGER = 536870912
	// BG_EVENT_USB_COMPLEX_TRIGGER as defined in beagle/beagle.h:846
	BG_EVENT_USB_COMPLEX_TRIGGER = 1073741824
	// BG_EVENT_USB_TRIGGER as defined in beagle/beagle.h:847
	BG_EVENT_USB_TRIGGER = 2147483648
	// BG_EVENT_USB_TRIGGER_STATE_MASK as defined in beagle/beagle.h:848
	BG_EVENT_USB_TRIGGER_STATE_MASK = 112
	// BG_EVENT_USB_TRIGGER_STATE_SHIFT as defined in beagle/beagle.h:849
	BG_EVENT_USB_TRIGGER_STATE_SHIFT = 4
	// BG_EVENT_USB_TRIGGER_STATE_0 as defined in beagle/beagle.h:850
	BG_EVENT_USB_TRIGGER_STATE_0 = 0
	// BG_EVENT_USB_TRIGGER_STATE_1 as defined in beagle/beagle.h:851
	BG_EVENT_USB_TRIGGER_STATE_1 = 16
	// BG_EVENT_USB_TRIGGER_STATE_2 as defined in beagle/beagle.h:852
	BG_EVENT_USB_TRIGGER_STATE_2 = 32
	// BG_EVENT_USB_TRIGGER_STATE_3 as defined in beagle/beagle.h:853
	BG_EVENT_USB_TRIGGER_STATE_3 = 48
	// BG_EVENT_USB_TRIGGER_STATE_4 as defined in beagle/beagle.h:854
	BG_EVENT_USB_TRIGGER_STATE_4 = 64
	// BG_EVENT_USB_TRIGGER_STATE_5 as defined in beagle/beagle.h:855
	BG_EVENT_USB_TRIGGER_STATE_5 = 80
	// BG_EVENT_USB_TRIGGER_STATE_6 as defined in beagle/beagle.h:856
	BG_EVENT_USB_TRIGGER_STATE_6 = 96
	// BG_EVENT_USB_TRIGGER_STATE_7 as defined in beagle/beagle.h:857
	BG_EVENT_USB_TRIGGER_STATE_7 = 112
	// BG_EVENT_USB_LFPS as defined in beagle/beagle.h:859
	BG_EVENT_USB_LFPS = 4096
	// BG_EVENT_USB_LTSSM as defined in beagle/beagle.h:860
	BG_EVENT_USB_LTSSM = 8192
	// BG_EVENT_USB_VBUS_PRESENT as defined in beagle/beagle.h:861
	BG_EVENT_USB_VBUS_PRESENT = 65536
	// BG_EVENT_USB_VBUS_ABSENT as defined in beagle/beagle.h:862
	BG_EVENT_USB_VBUS_ABSENT = 131072
	// BG_EVENT_USB_SCRAMBLING_ENABLED as defined in beagle/beagle.h:863
	BG_EVENT_USB_SCRAMBLING_ENABLED = 262144
	// BG_EVENT_USB_SCRAMBLING_DISABLED as defined in beagle/beagle.h:864
	BG_EVENT_USB_SCRAMBLING_DISABLED = 524288
	// BG_EVENT_USB_POLARITY_NORMAL as defined in beagle/beagle.h:865
	BG_EVENT_USB_POLARITY_NORMAL = 1048576
	// BG_EVENT_USB_POLARITY_REVERSED as defined in beagle/beagle.h:866
	BG_EVENT_USB_POLARITY_REVERSED = 2097152
	// BG_EVENT_USB_PHY_ERROR as defined in beagle/beagle.h:867
	BG_EVENT_USB_PHY_ERROR = 4194304
	// BG_EVENT_USB_LTSSM_MASK as defined in beagle/beagle.h:868
	BG_EVENT_USB_LTSSM_MASK = 255
	// BG_EVENT_USB_LTSSM_STATE_UNKNOWN as defined in beagle/beagle.h:869
	BG_EVENT_USB_LTSSM_STATE_UNKNOWN = 0
	// BG_EVENT_USB_LTSSM_STATE_SS_DISABLED as defined in beagle/beagle.h:870
	BG_EVENT_USB_LTSSM_STATE_SS_DISABLED = 1
	// BG_EVENT_USB_LTSSM_STATE_SS_INACTIVE as defined in beagle/beagle.h:871
	BG_EVENT_USB_LTSSM_STATE_SS_INACTIVE = 2
	// BG_EVENT_USB_LTSSM_STATE_RX_DETECT_RESET as defined in beagle/beagle.h:872
	BG_EVENT_USB_LTSSM_STATE_RX_DETECT_RESET = 3
	// BG_EVENT_USB_LTSSM_STATE_RX_DETECT_ACTIVE as defined in beagle/beagle.h:873
	BG_EVENT_USB_LTSSM_STATE_RX_DETECT_ACTIVE = 4
	// BG_EVENT_USB_LTSSM_STATE_POLLING_LFPS as defined in beagle/beagle.h:874
	BG_EVENT_USB_LTSSM_STATE_POLLING_LFPS = 5
	// BG_EVENT_USB_LTSSM_STATE_POLLING_RXEQ as defined in beagle/beagle.h:875
	BG_EVENT_USB_LTSSM_STATE_POLLING_RXEQ = 6
	// BG_EVENT_USB_LTSSM_STATE_POLLING_ACTIVE as defined in beagle/beagle.h:876
	BG_EVENT_USB_LTSSM_STATE_POLLING_ACTIVE = 7
	// BG_EVENT_USB_LTSSM_STATE_POLLING_CONFIG as defined in beagle/beagle.h:877
	BG_EVENT_USB_LTSSM_STATE_POLLING_CONFIG = 8
	// BG_EVENT_USB_LTSSM_STATE_POLLING_IDLE as defined in beagle/beagle.h:878
	BG_EVENT_USB_LTSSM_STATE_POLLING_IDLE = 9
	// BG_EVENT_USB_LTSSM_STATE_U0 as defined in beagle/beagle.h:879
	BG_EVENT_USB_LTSSM_STATE_U0 = 10
	// BG_EVENT_USB_LTSSM_STATE_U1 as defined in beagle/beagle.h:880
	BG_EVENT_USB_LTSSM_STATE_U1 = 11
	// BG_EVENT_USB_LTSSM_STATE_U2 as defined in beagle/beagle.h:881
	BG_EVENT_USB_LTSSM_STATE_U2 = 12
	// BG_EVENT_USB_LTSSM_STATE_U3 as defined in beagle/beagle.h:882
	BG_EVENT_USB_LTSSM_STATE_U3 = 13
	// BG_EVENT_USB_LTSSM_STATE_RECOVERY_ACTIVE as defined in beagle/beagle.h:883
	BG_EVENT_USB_LTSSM_STATE_RECOVERY_ACTIVE = 14
	// BG_EVENT_USB_LTSSM_STATE_RECOVERY_CONFIG as defined in beagle/beagle.h:884
	BG_EVENT_USB_LTSSM_STATE_RECOVERY_CONFIG = 15
	// BG_EVENT_USB_LTSSM_STATE_RECOVERY_IDLE as defined in beagle/beagle.h:885
	BG_EVENT_USB_LTSSM_STATE_RECOVERY_IDLE = 16
	// BG_EVENT_USB_LTSSM_STATE_HOT_RESET_ACTIVE as defined in beagle/beagle.h:886
	BG_EVENT_USB_LTSSM_STATE_HOT_RESET_ACTIVE = 17
	// BG_EVENT_USB_LTSSM_STATE_HOT_RESET_EXIT as defined in beagle/beagle.h:887
	BG_EVENT_USB_LTSSM_STATE_HOT_RESET_EXIT = 18
	// BG_EVENT_USB_LTSSM_STATE_LOOPBACK_ACTIVE as defined in beagle/beagle.h:888
	BG_EVENT_USB_LTSSM_STATE_LOOPBACK_ACTIVE = 19
	// BG_EVENT_USB_LTSSM_STATE_LOOPBACK_EXIT as defined in beagle/beagle.h:889
	BG_EVENT_USB_LTSSM_STATE_LOOPBACK_EXIT = 20
	// BG_EVENT_USB_LTSSM_STATE_COMPLIANCE as defined in beagle/beagle.h:890
	BG_EVENT_USB_LTSSM_STATE_COMPLIANCE = 21
	// BG_EVENT_USB_SMA_EXTIN_ASSERTED as defined in beagle/beagle.h:891
	BG_EVENT_USB_SMA_EXTIN_ASSERTED = 16777216
	// BG_EVENT_USB_SMA_EXTIN_DEASSERTED as defined in beagle/beagle.h:892
	BG_EVENT_USB_SMA_EXTIN_DEASSERTED = 33554432
	// BG_EVENT_USB_TRIGGER_5GBIT_START as defined in beagle/beagle.h:893
	BG_EVENT_USB_TRIGGER_5GBIT_START = 67108864
	// BG_EVENT_USB_TRIGGER_5GBIT_STOP as defined in beagle/beagle.h:894
	BG_EVENT_USB_TRIGGER_5GBIT_STOP = 134217728
	// BG_EVENT_USBPD_CC_MASK as defined in beagle/beagle.h:896
	BG_EVENT_USBPD_CC_MASK = 4026531840
	// BG_EVENT_USBPD_CC_SHIFT as defined in beagle/beagle.h:897
	BG_EVENT_USBPD_CC_SHIFT = 28
	// BG_EVENT_USBPD_CC1 as defined in beagle/beagle.h:898
	BG_EVENT_USBPD_CC1 = 0
	// BG_EVENT_USBPD_CC2 as defined in beagle/beagle.h:899
	BG_EVENT_USBPD_CC2 = 268435456
	// BG_EVENT_USBPD_POL_CHANGE as defined in beagle/beagle.h:900
	BG_EVENT_USBPD_POL_CHANGE = 8
	// BG_EVENT_USBPD_SOP_MASK as defined in beagle/beagle.h:901
	BG_EVENT_USBPD_SOP_MASK = 7
	// BG_EVENT_USBPD_SOP as defined in beagle/beagle.h:902
	BG_EVENT_USBPD_SOP = 0
	// BG_EVENT_USBPD_SOP_PRIME as defined in beagle/beagle.h:903
	BG_EVENT_USBPD_SOP_PRIME = 1
	// BG_EVENT_USBPD_SOP_DPRIME as defined in beagle/beagle.h:904
	BG_EVENT_USBPD_SOP_DPRIME = 2
	// BG_EVENT_USBPD_SOP_PRIME_DEBUG as defined in beagle/beagle.h:905
	BG_EVENT_USBPD_SOP_PRIME_DEBUG = 3
	// BG_EVENT_USBPD_SOP_DPRIME_DEBUG as defined in beagle/beagle.h:906
	BG_EVENT_USBPD_SOP_DPRIME_DEBUG = 4
	// BG_EVENT_USBPD_HARD_RESET as defined in beagle/beagle.h:907
	BG_EVENT_USBPD_HARD_RESET = 6
	// BG_EVENT_USBPD_CABLE_RESET as defined in beagle/beagle.h:908
	BG_EVENT_USBPD_CABLE_RESET = 7
	// BG_EVENT_USBPD_EXTENDED_HEADER as defined in beagle/beagle.h:909
	BG_EVENT_USBPD_EXTENDED_HEADER = 16
	// BG_USB_FEATURE_NONE as defined in beagle/beagle.h:911
	BG_USB_FEATURE_NONE = 0
	// BG_USB_FEATURE_USB2_MON as defined in beagle/beagle.h:912
	BG_USB_FEATURE_USB2_MON = 1
	// BG_USB_FEATURE_USB3_MON as defined in beagle/beagle.h:913
	BG_USB_FEATURE_USB3_MON = 2
	// BG_USB_FEATURE_SIMUL_23 as defined in beagle/beagle.h:914
	BG_USB_FEATURE_SIMUL_23 = 4
	// BG_USB_FEATURE_USB3_CMP_TRIG as defined in beagle/beagle.h:915
	BG_USB_FEATURE_USB3_CMP_TRIG = 8
	// BG_USB_FEATURE_USB3_4G_MEM as defined in beagle/beagle.h:916
	BG_USB_FEATURE_USB3_4G_MEM = 32
	// BG_USB_FEATURE_USB2_CMP_TRIG as defined in beagle/beagle.h:917
	BG_USB_FEATURE_USB2_CMP_TRIG = 128
	// BG_USB_FEATURE_CROSS_ANALYZER_SYNC as defined in beagle/beagle.h:918
	BG_USB_FEATURE_CROSS_ANALYZER_SYNC = 256
	// BG_USB_FEATURE_USB3_DOWNLINK as defined in beagle/beagle.h:919
	BG_USB_FEATURE_USB3_DOWNLINK = 512
	// BG_USB_FEATURE_IV_MON_LITE as defined in beagle/beagle.h:920
	BG_USB_FEATURE_IV_MON_LITE = 1024
	// BG_USB_FEATURE_USBPD as defined in beagle/beagle.h:921
	BG_USB_FEATURE_USBPD = 2048
	// BG_USB_LICENSE_LENGTH as defined in beagle/beagle.h:928
	BG_USB_LICENSE_LENGTH = 60
	// BG_USB_CAPTURE_USB3 as defined in beagle/beagle.h:956
	BG_USB_CAPTURE_USB3 = 1
	// BG_USB_CAPTURE_USB2 as defined in beagle/beagle.h:957
	BG_USB_CAPTURE_USB2 = 2
	// BG_USB_CAPTURE_IV_MON_LITE as defined in beagle/beagle.h:958
	BG_USB_CAPTURE_IV_MON_LITE = 8
	// BG_USB_CAPTURE_USBPD as defined in beagle/beagle.h:959
	BG_USB_CAPTURE_USBPD = 16
	// BG_USB2_AUTO_SPEED_DETECT as defined in beagle/beagle.h:1009
	BG_USB2_AUTO_SPEED_DETECT = 0
	// BG_USB2_LOW_SPEED as defined in beagle/beagle.h:1010
	BG_USB2_LOW_SPEED = 1
	// BG_USB2_FULL_SPEED as defined in beagle/beagle.h:1011
	BG_USB2_FULL_SPEED = 2
	// BG_USB2_HIGH_SPEED as defined in beagle/beagle.h:1012
	BG_USB2_HIGH_SPEED = 3
	// BG_USB2_VBUS_OVERRIDE as defined in beagle/beagle.h:1013
	BG_USB2_VBUS_OVERRIDE = 128
	// BG_USB_CAPTURE_SIZE_INFINITE as defined in beagle/beagle.h:1021
	BG_USB_CAPTURE_SIZE_INFINITE = 0
	// BG_USB_CAPTURE_SIZE_SCALE as defined in beagle/beagle.h:1022
	BG_USB_CAPTURE_SIZE_SCALE = 4294967295
	// BG_USB2_DIGITAL_OUT_ENABLE_PIN1 as defined in beagle/beagle.h:1049
	BG_USB2_DIGITAL_OUT_ENABLE_PIN1 = 1
	// BG_USB2_DIGITAL_OUT_PIN1_ACTIVE_HIGH as defined in beagle/beagle.h:1050
	BG_USB2_DIGITAL_OUT_PIN1_ACTIVE_HIGH = 1
	// BG_USB2_DIGITAL_OUT_PIN1_ACTIVE_LOW as defined in beagle/beagle.h:1051
	BG_USB2_DIGITAL_OUT_PIN1_ACTIVE_LOW = 0
	// BG_USB2_DIGITAL_OUT_ENABLE_PIN2 as defined in beagle/beagle.h:1052
	BG_USB2_DIGITAL_OUT_ENABLE_PIN2 = 2
	// BG_USB2_DIGITAL_OUT_PIN2_ACTIVE_HIGH as defined in beagle/beagle.h:1053
	BG_USB2_DIGITAL_OUT_PIN2_ACTIVE_HIGH = 2
	// BG_USB2_DIGITAL_OUT_PIN2_ACTIVE_LOW as defined in beagle/beagle.h:1054
	BG_USB2_DIGITAL_OUT_PIN2_ACTIVE_LOW = 0
	// BG_USB2_DIGITAL_OUT_ENABLE_PIN3 as defined in beagle/beagle.h:1055
	BG_USB2_DIGITAL_OUT_ENABLE_PIN3 = 4
	// BG_USB2_DIGITAL_OUT_PIN3_ACTIVE_HIGH as defined in beagle/beagle.h:1056
	BG_USB2_DIGITAL_OUT_PIN3_ACTIVE_HIGH = 4
	// BG_USB2_DIGITAL_OUT_PIN3_ACTIVE_LOW as defined in beagle/beagle.h:1057
	BG_USB2_DIGITAL_OUT_PIN3_ACTIVE_LOW = 0
	// BG_USB2_DIGITAL_OUT_ENABLE_PIN4 as defined in beagle/beagle.h:1058
	BG_USB2_DIGITAL_OUT_ENABLE_PIN4 = 8
	// BG_USB2_DIGITAL_OUT_PIN4_ACTIVE_HIGH as defined in beagle/beagle.h:1059
	BG_USB2_DIGITAL_OUT_PIN4_ACTIVE_HIGH = 8
	// BG_USB2_DIGITAL_OUT_PIN4_ACTIVE_LOW as defined in beagle/beagle.h:1060
	BG_USB2_DIGITAL_OUT_PIN4_ACTIVE_LOW = 0
	// BG_USB2_DATA_MATCH_DATA0 as defined in beagle/beagle.h:1101
	BG_USB2_DATA_MATCH_DATA0 = 1
	// BG_USB2_DATA_MATCH_DATA1 as defined in beagle/beagle.h:1102
	BG_USB2_DATA_MATCH_DATA1 = 2
	// BG_USB2_DATA_MATCH_DATA2 as defined in beagle/beagle.h:1103
	BG_USB2_DATA_MATCH_DATA2 = 4
	// BG_USB2_DATA_MATCH_MDATA as defined in beagle/beagle.h:1104
	BG_USB2_DATA_MATCH_MDATA = 8
	// BG_USB2_DIGITAL_IN_ENABLE_PIN1 as defined in beagle/beagle.h:1126
	BG_USB2_DIGITAL_IN_ENABLE_PIN1 = 1
	// BG_USB2_DIGITAL_IN_ENABLE_PIN2 as defined in beagle/beagle.h:1127
	BG_USB2_DIGITAL_IN_ENABLE_PIN2 = 2
	// BG_USB2_DIGITAL_IN_ENABLE_PIN3 as defined in beagle/beagle.h:1128
	BG_USB2_DIGITAL_IN_ENABLE_PIN3 = 4
	// BG_USB2_DIGITAL_IN_ENABLE_PIN4 as defined in beagle/beagle.h:1129
	BG_USB2_DIGITAL_IN_ENABLE_PIN4 = 8
	// BG_USB2_HW_FILTER_PID_SOF as defined in beagle/beagle.h:1137
	BG_USB2_HW_FILTER_PID_SOF = 1
	// BG_USB2_HW_FILTER_PID_IN as defined in beagle/beagle.h:1138
	BG_USB2_HW_FILTER_PID_IN = 2
	// BG_USB2_HW_FILTER_PID_PING as defined in beagle/beagle.h:1139
	BG_USB2_HW_FILTER_PID_PING = 4
	// BG_USB2_HW_FILTER_PID_PRE as defined in beagle/beagle.h:1140
	BG_USB2_HW_FILTER_PID_PRE = 8
	// BG_USB2_HW_FILTER_PID_SPLIT as defined in beagle/beagle.h:1141
	BG_USB2_HW_FILTER_PID_SPLIT = 16
	// BG_USB2_HW_FILTER_SELF as defined in beagle/beagle.h:1142
	BG_USB2_HW_FILTER_SELF = 32
	// BG_USB2_MATCH_HANDSHAKE_MASK_DISABLED as defined in beagle/beagle.h:1247
	BG_USB2_MATCH_HANDSHAKE_MASK_DISABLED = 0
	// BG_USB2_MATCH_HANDSHAKE_MASK_NONE as defined in beagle/beagle.h:1248
	BG_USB2_MATCH_HANDSHAKE_MASK_NONE = 1
	// BG_USB2_MATCH_HANDSHAKE_MASK_ACK as defined in beagle/beagle.h:1249
	BG_USB2_MATCH_HANDSHAKE_MASK_ACK = 2
	// BG_USB2_MATCH_HANDSHAKE_MASK_NAK as defined in beagle/beagle.h:1250
	BG_USB2_MATCH_HANDSHAKE_MASK_NAK = 4
	// BG_USB2_MATCH_HANDSHAKE_MASK_NYET as defined in beagle/beagle.h:1251
	BG_USB2_MATCH_HANDSHAKE_MASK_NYET = 8
	// BG_USB2_MATCH_HANDSHAKE_MASK_STALL as defined in beagle/beagle.h:1252
	BG_USB2_MATCH_HANDSHAKE_MASK_STALL = 16
	// BG_USB_COMPLEX_MATCH_ACTION_EXTOUT as defined in beagle/beagle.h:1276
	BG_USB_COMPLEX_MATCH_ACTION_EXTOUT = 1
	// BG_USB_COMPLEX_MATCH_ACTION_TRIGGER as defined in beagle/beagle.h:1277
	BG_USB_COMPLEX_MATCH_ACTION_TRIGGER = 2
	// BG_USB_COMPLEX_MATCH_ACTION_FILTER as defined in beagle/beagle.h:1278
	BG_USB_COMPLEX_MATCH_ACTION_FILTER = 4
	// BG_USB_COMPLEX_MATCH_ACTION_GOTO as defined in beagle/beagle.h:1279
	BG_USB_COMPLEX_MATCH_ACTION_GOTO = 8
	// BG_USB3_PHY_CONFIG_POLARITY_NON_INVERT as defined in beagle/beagle.h:1462
	BG_USB3_PHY_CONFIG_POLARITY_NON_INVERT = 0
	// BG_USB3_PHY_CONFIG_POLARITY_INVERT as defined in beagle/beagle.h:1463
	BG_USB3_PHY_CONFIG_POLARITY_INVERT = 1
	// BG_USB3_PHY_CONFIG_POLARITY_AUTO as defined in beagle/beagle.h:1464
	BG_USB3_PHY_CONFIG_POLARITY_AUTO = 2
	// BG_USB3_PHY_CONFIG_POLARITY_MASK as defined in beagle/beagle.h:1465
	BG_USB3_PHY_CONFIG_POLARITY_MASK = 3
	// BG_USB3_PHY_CONFIG_DESCRAMBLER_ON as defined in beagle/beagle.h:1466
	BG_USB3_PHY_CONFIG_DESCRAMBLER_ON = 0
	// BG_USB3_PHY_CONFIG_DESCRAMBLER_OFF as defined in beagle/beagle.h:1467
	BG_USB3_PHY_CONFIG_DESCRAMBLER_OFF = 4
	// BG_USB3_PHY_CONFIG_DESCRAMBLER_AUTO as defined in beagle/beagle.h:1468
	BG_USB3_PHY_CONFIG_DESCRAMBLER_AUTO = 8
	// BG_USB3_PHY_CONFIG_DESCRAMBLER_MASK as defined in beagle/beagle.h:1469
	BG_USB3_PHY_CONFIG_DESCRAMBLER_MASK = 12
	// BG_USB3_PHY_CONFIG_RXTERM_ON as defined in beagle/beagle.h:1470
	BG_USB3_PHY_CONFIG_RXTERM_ON = 0
	// BG_USB3_PHY_CONFIG_RXTERM_OFF as defined in beagle/beagle.h:1471
	BG_USB3_PHY_CONFIG_RXTERM_OFF = 16
	// BG_USB3_PHY_CONFIG_RXTERM_AUTO as defined in beagle/beagle.h:1472
	BG_USB3_PHY_CONFIG_RXTERM_AUTO = 32
	// BG_USB3_PHY_CONFIG_RXTERM_MASK as defined in beagle/beagle.h:1473
	BG_USB3_PHY_CONFIG_RXTERM_MASK = 48
	// BG_USB3_TRUNCATION_OFF as defined in beagle/beagle.h:1481
	BG_USB3_TRUNCATION_OFF = 0
	// BG_USB3_TRUNCATION_20 as defined in beagle/beagle.h:1482
	BG_USB3_TRUNCATION_20 = 1
	// BG_USB3_TRUNCATION_36 as defined in beagle/beagle.h:1483
	BG_USB3_TRUNCATION_36 = 2
	// BG_USB3_TRUNCATION_68 as defined in beagle/beagle.h:1484
	BG_USB3_TRUNCATION_68 = 3
	// BG_USB3_EQUALIZATION_OFF as defined in beagle/beagle.h:1493
	BG_USB3_EQUALIZATION_OFF = 0
	// BG_USB3_EQUALIZATION_MIN as defined in beagle/beagle.h:1494
	BG_USB3_EQUALIZATION_MIN = 1
	// BG_USB3_EQUALIZATION_MOD as defined in beagle/beagle.h:1495
	BG_USB3_EQUALIZATION_MOD = 2
	// BG_USB3_EQUALIZATION_MAX as defined in beagle/beagle.h:1496
	BG_USB3_EQUALIZATION_MAX = 3
	// BG_USB3_SIMPLE_MATCH_NONE as defined in beagle/beagle.h:1519
	BG_USB3_SIMPLE_MATCH_NONE = 0
	// BG_USB3_SIMPLE_MATCH_SSTX_IPS as defined in beagle/beagle.h:1520
	BG_USB3_SIMPLE_MATCH_SSTX_IPS = 1
	// BG_USB3_SIMPLE_MATCH_SSTX_SLC as defined in beagle/beagle.h:1521
	BG_USB3_SIMPLE_MATCH_SSTX_SLC = 2
	// BG_USB3_SIMPLE_MATCH_SSTX_SHP as defined in beagle/beagle.h:1522
	BG_USB3_SIMPLE_MATCH_SSTX_SHP = 4
	// BG_USB3_SIMPLE_MATCH_SSTX_SDP as defined in beagle/beagle.h:1523
	BG_USB3_SIMPLE_MATCH_SSTX_SDP = 8
	// BG_USB3_SIMPLE_MATCH_SSRX_IPS as defined in beagle/beagle.h:1524
	BG_USB3_SIMPLE_MATCH_SSRX_IPS = 16
	// BG_USB3_SIMPLE_MATCH_SSRX_SLC as defined in beagle/beagle.h:1525
	BG_USB3_SIMPLE_MATCH_SSRX_SLC = 32
	// BG_USB3_SIMPLE_MATCH_SSRX_SHP as defined in beagle/beagle.h:1526
	BG_USB3_SIMPLE_MATCH_SSRX_SHP = 64
	// BG_USB3_SIMPLE_MATCH_SSRX_SDP as defined in beagle/beagle.h:1527
	BG_USB3_SIMPLE_MATCH_SSRX_SDP = 128
	// BG_USB3_SIMPLE_MATCH_SSTX_SLC_CRC_5A_CRC_5B as defined in beagle/beagle.h:1528
	BG_USB3_SIMPLE_MATCH_SSTX_SLC_CRC_5A_CRC_5B = 256
	// BG_USB3_SIMPLE_MATCH_SSTX_SHP_CRC_5 as defined in beagle/beagle.h:1529
	BG_USB3_SIMPLE_MATCH_SSTX_SHP_CRC_5 = 512
	// BG_USB3_SIMPLE_MATCH_SSTX_SHP_CRC_16 as defined in beagle/beagle.h:1530
	BG_USB3_SIMPLE_MATCH_SSTX_SHP_CRC_16 = 1024
	// BG_USB3_SIMPLE_MATCH_SSTX_SDP_CRC as defined in beagle/beagle.h:1531
	BG_USB3_SIMPLE_MATCH_SSTX_SDP_CRC = 2048
	// BG_USB3_SIMPLE_MATCH_SSRX_SLC_CRC_5A_CRC_5B as defined in beagle/beagle.h:1532
	BG_USB3_SIMPLE_MATCH_SSRX_SLC_CRC_5A_CRC_5B = 4096
	// BG_USB3_SIMPLE_MATCH_SSRX_SHP_CRC_5 as defined in beagle/beagle.h:1533
	BG_USB3_SIMPLE_MATCH_SSRX_SHP_CRC_5 = 8192
	// BG_USB3_SIMPLE_MATCH_SSRX_SHP_CRC_16 as defined in beagle/beagle.h:1534
	BG_USB3_SIMPLE_MATCH_SSRX_SHP_CRC_16 = 16384
	// BG_USB3_SIMPLE_MATCH_SSRX_SDP_CRC as defined in beagle/beagle.h:1535
	BG_USB3_SIMPLE_MATCH_SSRX_SDP_CRC = 32768
	// BG_USB3_SIMPLE_MATCH_EVENT_SSTX_LFPS as defined in beagle/beagle.h:1536
	BG_USB3_SIMPLE_MATCH_EVENT_SSTX_LFPS = 65536
	// BG_USB3_SIMPLE_MATCH_EVENT_SSTX_POLARITY as defined in beagle/beagle.h:1537
	BG_USB3_SIMPLE_MATCH_EVENT_SSTX_POLARITY = 131072
	// BG_USB3_SIMPLE_MATCH_EVENT_SSTX_DETECTED as defined in beagle/beagle.h:1538
	BG_USB3_SIMPLE_MATCH_EVENT_SSTX_DETECTED = 4194304
	// BG_USB3_SIMPLE_MATCH_EVENT_SSTX_SCRAMBLE as defined in beagle/beagle.h:1539
	BG_USB3_SIMPLE_MATCH_EVENT_SSTX_SCRAMBLE = 524288
	// BG_USB3_SIMPLE_MATCH_EVENT_SSRX_LFPS as defined in beagle/beagle.h:1540
	BG_USB3_SIMPLE_MATCH_EVENT_SSRX_LFPS = 1048576
	// BG_USB3_SIMPLE_MATCH_EVENT_SSRX_POLARITY as defined in beagle/beagle.h:1541
	BG_USB3_SIMPLE_MATCH_EVENT_SSRX_POLARITY = 2097152
	// BG_USB3_SIMPLE_MATCH_EVENT_SSRX_DETECTED as defined in beagle/beagle.h:1542
	BG_USB3_SIMPLE_MATCH_EVENT_SSRX_DETECTED = 262144
	// BG_USB3_SIMPLE_MATCH_EVENT_SSRX_SCRAMBLE as defined in beagle/beagle.h:1543
	BG_USB3_SIMPLE_MATCH_EVENT_SSRX_SCRAMBLE = 8388608
	// BG_USB3_SIMPLE_MATCH_EVENT_VBUS_PRESENT as defined in beagle/beagle.h:1544
	BG_USB3_SIMPLE_MATCH_EVENT_VBUS_PRESENT = 134217728
	// BG_USB3_SIMPLE_MATCH_EVENT_SSTX_PHYERR as defined in beagle/beagle.h:1545
	BG_USB3_SIMPLE_MATCH_EVENT_SSTX_PHYERR = 268435456
	// BG_USB3_SIMPLE_MATCH_EVENT_SSRX_PHYERR as defined in beagle/beagle.h:1546
	BG_USB3_SIMPLE_MATCH_EVENT_SSRX_PHYERR = 536870912
	// BG_USB3_SIMPLE_MATCH_EVENT_SMA_EXTIN as defined in beagle/beagle.h:1547
	BG_USB3_SIMPLE_MATCH_EVENT_SMA_EXTIN = 1073741824
	// BG_USB_EDGE_RISING as defined in beagle/beagle.h:1548
	BG_USB_EDGE_RISING = 1
	// BG_USB_EDGE_PULSE as defined in beagle/beagle.h:1549
	BG_USB_EDGE_PULSE = 1
	// BG_USB_EDGE_FALLING as defined in beagle/beagle.h:1550
	BG_USB_EDGE_FALLING = 2
	// BG_USB_EDGE_DEVICE_CHIRP as defined in beagle/beagle.h:1551
	BG_USB_EDGE_DEVICE_CHIRP = 1
	// BG_USB_EDGE_HOST_CHIRP as defined in beagle/beagle.h:1552
	BG_USB_EDGE_HOST_CHIRP = 2
	// BG480_TRUNCATION_LENGTH as defined in beagle/beagle.h:1977
	BG480_TRUNCATION_LENGTH = 4
	// BG480V2_USB2_BUFFER_SIZE_256MB as defined in beagle/beagle.h:1978
	BG480V2_USB2_BUFFER_SIZE_256MB = 256
	// BG5000_USB2_BUFFER_SIZE_128MB as defined in beagle/beagle.h:1984
	BG5000_USB2_BUFFER_SIZE_128MB = 128
	// BG5000_USB3_BUFFER_SIZE_2GB as defined in beagle/beagle.h:1985
	BG5000_USB3_BUFFER_SIZE_2GB = 2
	// BG5000_USB3_BUFFER_SIZE_4GB as defined in beagle/beagle.h:1986
	BG5000_USB3_BUFFER_SIZE_4GB = 4
	// BG_MDIO_OPCODE22_WRITE as defined in beagle/beagle.h:2030
	BG_MDIO_OPCODE22_WRITE = 1
	// BG_MDIO_OPCODE22_READ as defined in beagle/beagle.h:2031
	BG_MDIO_OPCODE22_READ = 2
	// BG_MDIO_OPCODE22_ERROR as defined in beagle/beagle.h:2032
	BG_MDIO_OPCODE22_ERROR = 255
	// BG_MDIO_OPCODE45_ADDR as defined in beagle/beagle.h:2033
	BG_MDIO_OPCODE45_ADDR = 0
	// BG_MDIO_OPCODE45_WRITE as defined in beagle/beagle.h:2034
	BG_MDIO_OPCODE45_WRITE = 1
	// BG_MDIO_OPCODE45_READ_POSTINC as defined in beagle/beagle.h:2035
	BG_MDIO_OPCODE45_READ_POSTINC = 2
	// BG_MDIO_OPCODE45_READ as defined in beagle/beagle.h:2036
	BG_MDIO_OPCODE45_READ = 3
	// BG_HEADER_USBPD_STANDARD_MASK as defined in beagle/beagle.h:2096
	BG_HEADER_USBPD_STANDARD_MASK = 65535
	// BG_HEADER_USBPD_EXTENDED_MASK as defined in beagle/beagle.h:2097
	BG_HEADER_USBPD_EXTENDED_MASK = 4294901760
)
