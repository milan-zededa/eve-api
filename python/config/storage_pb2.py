# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: config/storage.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from config import devcommon_pb2 as config_dot_devcommon__pb2
from evecommon import acipherinfo_pb2 as evecommon_dot_acipherinfo__pb2
from evecommon import evecommon_pb2 as evecommon_dot_evecommon__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x14\x63onfig/storage.proto\x12\x15org.lfedge.eve.config\x1a\x16\x63onfig/devcommon.proto\x1a\x1b\x65vecommon/acipherinfo.proto\x1a\x19\x65vecommon/evecommon.proto\"P\n\rSignatureInfo\x12\x15\n\rintercertsurl\x18\x01 \x01(\t\x12\x15\n\rsignercerturl\x18\x02 \x01(\t\x12\x11\n\tsignature\x18\x03 \x01(\x0c\"\xe5\x01\n\x0f\x44\x61tastoreConfig\x12\n\n\x02id\x18\x64 \x01(\t\x12,\n\x05\x64Type\x18\x01 \x01(\x0e\x32\x1d.org.lfedge.eve.config.DsType\x12\x0c\n\x04\x66qdn\x18\x02 \x01(\t\x12\x0e\n\x06\x61piKey\x18\x03 \x01(\t\x12\x10\n\x08password\x18\x04 \x01(\t\x12\r\n\x05\x64path\x18\x05 \x01(\t\x12\x0e\n\x06region\x18\x06 \x01(\t\x12\x36\n\ncipherData\x18\x07 \x01(\x0b\x32\".org.lfedge.eve.common.CipherBlock\x12\x11\n\tdsCertPEM\x18\x08 \x03(\x0c\"\xec\x01\n\x05Image\x12=\n\x0euuidandversion\x18\x01 \x01(\x0b\x32%.org.lfedge.eve.config.UUIDandVersion\x12\x0c\n\x04name\x18\x02 \x01(\t\x12\x0e\n\x06sha256\x18\x03 \x01(\t\x12.\n\x07iformat\x18\x04 \x01(\x0e\x32\x1d.org.lfedge.eve.config.Format\x12\x35\n\x07siginfo\x18\x05 \x01(\x0b\x32$.org.lfedge.eve.config.SignatureInfo\x12\x0c\n\x04\x64sId\x18\x06 \x01(\t\x12\x11\n\tsizeBytes\x18\x08 \x01(\x03\"\xd0\x01\n\x05\x44rive\x12+\n\x05image\x18\x01 \x01(\x0b\x32\x1c.org.lfedge.eve.config.Image\x12\x10\n\x08readonly\x18\x05 \x01(\x08\x12\x10\n\x08preserve\x18\x06 \x01(\x08\x12\x31\n\x07\x64rvtype\x18\x08 \x01(\x0e\x32 .org.lfedge.eve.config.DriveType\x12-\n\x06target\x18\t \x01(\x0e\x32\x1d.org.lfedge.eve.config.Target\x12\x14\n\x0cmaxsizebytes\x18\n \x01(\x03\"\xa1\x02\n\x0b\x43ontentTree\x12\x0c\n\x04uuid\x18\x01 \x01(\t\x12\x0c\n\x04\x64sId\x18\x02 \x01(\t\x12\x0b\n\x03URL\x18\x03 \x01(\t\x12.\n\x07iformat\x18\x04 \x01(\x0e\x32\x1d.org.lfedge.eve.config.Format\x12\x0e\n\x06sha256\x18\x05 \x01(\t\x12\x14\n\x0cmaxSizeBytes\x18\x06 \x01(\x04\x12\x35\n\x07siginfo\x18\x07 \x01(\x0b\x32$.org.lfedge.eve.config.SignatureInfo\x12\x13\n\x0b\x64isplayName\x18\x08 \x01(\t\x12\x18\n\x10generation_count\x18\t \x01(\x03\x12\x18\n\x10\x63ustom_meta_data\x18\n \x01(\t\x12\x13\n\x0b\x64s_ids_list\x18\x0b \x03(\t\"r\n\x13VolumeContentOrigin\x12<\n\x04type\x18\x01 \x01(\x0e\x32..org.lfedge.eve.config.VolumeContentOriginType\x12\x1d\n\x15\x64ownloadContentTreeID\x18\x02 \x01(\t\"\xc8\x02\n\x06Volume\x12\x0c\n\x04uuid\x18\x01 \x01(\t\x12:\n\x06origin\x18\x02 \x01(\x0b\x32*.org.lfedge.eve.config.VolumeContentOrigin\x12?\n\tprotocols\x18\x03 \x03(\x0e\x32,.org.lfedge.eve.config.VolumeAccessProtocols\x12\x17\n\x0fgenerationCount\x18\x04 \x01(\x03\x12\x14\n\x0cmaxsizebytes\x18\x05 \x01(\x03\x12\x10\n\x08readonly\x18\x06 \x01(\x08\x12\x13\n\x0b\x64isplayName\x18\x07 \x01(\t\x12\x12\n\nclear_text\x18\x08 \x01(\x08\x12-\n\x06target\x18\t \x01(\x0e\x32\x1d.org.lfedge.eve.config.Target\x12\x1a\n\x12\x64\x65signated_node_id\x18\n \x01(\t\"\xb8\x01\n\nDiskConfig\x12\x34\n\x04\x64isk\x18\x01 \x01(\x0b\x32&.org.lfedge.eve.common.DiskDescription\x12\x38\n\x08old_disk\x18\x02 \x01(\x0b\x32&.org.lfedge.eve.common.DiskDescription\x12:\n\x0b\x64isk_config\x18\x03 \x01(\x0e\x32%.org.lfedge.eve.config.DiskConfigType\"\xb0\x01\n\x0b\x44isksConfig\x12\x30\n\x05\x64isks\x18\x01 \x03(\x0b\x32!.org.lfedge.eve.config.DiskConfig\x12\x39\n\narray_type\x18\x02 \x01(\x0e\x32%.org.lfedge.eve.config.DisksArrayType\x12\x34\n\x08\x63hildren\x18\x03 \x03(\x0b\x32\".org.lfedge.eve.config.DisksConfig*\x85\x01\n\x06\x44sType\x12\r\n\tDsUnknown\x10\x00\x12\n\n\x06\x44sHttp\x10\x01\x12\x0b\n\x07\x44sHttps\x10\x02\x12\x08\n\x04\x44sS3\x10\x03\x12\n\n\x06\x44sSFTP\x10\x04\x12\x17\n\x13\x44sContainerRegistry\x10\x05\x12\x0f\n\x0b\x44sAzureBlob\x10\x06\x12\x13\n\x0f\x44sGoogleStorage\x10\x07*}\n\x06\x46ormat\x12\x0e\n\nFmtUnknown\x10\x00\x12\x07\n\x03RAW\x10\x01\x12\x08\n\x04QCOW\x10\x02\x12\t\n\x05QCOW2\x10\x03\x12\x07\n\x03VHD\x10\x04\x12\x08\n\x04VMDK\x10\x05\x12\x07\n\x03OVA\x10\x06\x12\x08\n\x04VHDX\x10\x07\x12\r\n\tCONTAINER\x10\x08\x12\x07\n\x03ISO\x10\t\x12\x07\n\x03PVC\x10\n*V\n\x06Target\x12\x0e\n\nTgtUnknown\x10\x00\x12\x08\n\x04\x44isk\x10\x01\x12\n\n\x06Kernel\x10\x02\x12\n\n\x06Initrd\x10\x03\x12\x0b\n\x07RamDisk\x10\x04\x12\r\n\tAppCustom\x10\x05*I\n\tDriveType\x12\x10\n\x0cUnclassified\x10\x00\x12\t\n\x05\x43\x44ROM\x10\x01\x12\x07\n\x03HDD\x10\x02\x12\x07\n\x03NET\x10\x03\x12\r\n\tHDD_EMPTY\x10\x04*1\n\x15VolumeAccessProtocols\x12\x0c\n\x08VAP_NONE\x10\x00\x12\n\n\x06VAP_9P\x10\x01*N\n\x17VolumeContentOriginType\x12\x10\n\x0cVCOT_UNKNOWN\x10\x00\x12\x0e\n\nVCOT_BLANK\x10\x01\x12\x11\n\rVCOT_DOWNLOAD\x10\x02*\xec\x01\n\x0e\x44iskConfigType\x12 \n\x1c\x44ISK_CONFIG_TYPE_UNSPECIFIED\x10\x00\x12\x1a\n\x16\x44ISK_CONFIG_TYPE_EVEOS\x10\x01\x12\x1c\n\x18\x44ISK_CONFIG_TYPE_PERSIST\x10\x02\x12\x1f\n\x1b\x44ISK_CONFIG_TYPE_ZFS_ONLINE\x10\x03\x12 \n\x1c\x44ISK_CONFIG_TYPE_ZFS_OFFLINE\x10\x04\x12\x1e\n\x1a\x44ISK_CONFIG_TYPE_APPDIRECT\x10\x05\x12\x1b\n\x17\x44ISK_CONFIG_TYPE_UNUSED\x10\x06*\xa2\x01\n\x0e\x44isksArrayType\x12 \n\x1c\x44ISKS_ARRAY_TYPE_UNSPECIFIED\x10\x00\x12\x1a\n\x16\x44ISKS_ARRAY_TYPE_RAID0\x10\x01\x12\x1a\n\x16\x44ISKS_ARRAY_TYPE_RAID1\x10\x02\x12\x1a\n\x16\x44ISKS_ARRAY_TYPE_RAID5\x10\x03\x12\x1a\n\x16\x44ISKS_ARRAY_TYPE_RAID6\x10\x04\x42=\n\x15org.lfedge.eve.configZ$github.com/lf-edge/eve-api/go/configb\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'config.storage_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'\n\025org.lfedge.eve.configZ$github.com/lf-edge/eve-api/go/config'
  _globals['_DSTYPE']._serialized_start=1997
  _globals['_DSTYPE']._serialized_end=2130
  _globals['_FORMAT']._serialized_start=2132
  _globals['_FORMAT']._serialized_end=2257
  _globals['_TARGET']._serialized_start=2259
  _globals['_TARGET']._serialized_end=2345
  _globals['_DRIVETYPE']._serialized_start=2347
  _globals['_DRIVETYPE']._serialized_end=2420
  _globals['_VOLUMEACCESSPROTOCOLS']._serialized_start=2422
  _globals['_VOLUMEACCESSPROTOCOLS']._serialized_end=2471
  _globals['_VOLUMECONTENTORIGINTYPE']._serialized_start=2473
  _globals['_VOLUMECONTENTORIGINTYPE']._serialized_end=2551
  _globals['_DISKCONFIGTYPE']._serialized_start=2554
  _globals['_DISKCONFIGTYPE']._serialized_end=2790
  _globals['_DISKSARRAYTYPE']._serialized_start=2793
  _globals['_DISKSARRAYTYPE']._serialized_end=2955
  _globals['_SIGNATUREINFO']._serialized_start=127
  _globals['_SIGNATUREINFO']._serialized_end=207
  _globals['_DATASTORECONFIG']._serialized_start=210
  _globals['_DATASTORECONFIG']._serialized_end=439
  _globals['_IMAGE']._serialized_start=442
  _globals['_IMAGE']._serialized_end=678
  _globals['_DRIVE']._serialized_start=681
  _globals['_DRIVE']._serialized_end=889
  _globals['_CONTENTTREE']._serialized_start=892
  _globals['_CONTENTTREE']._serialized_end=1181
  _globals['_VOLUMECONTENTORIGIN']._serialized_start=1183
  _globals['_VOLUMECONTENTORIGIN']._serialized_end=1297
  _globals['_VOLUME']._serialized_start=1300
  _globals['_VOLUME']._serialized_end=1628
  _globals['_DISKCONFIG']._serialized_start=1631
  _globals['_DISKCONFIG']._serialized_end=1815
  _globals['_DISKSCONFIG']._serialized_start=1818
  _globals['_DISKSCONFIG']._serialized_end=1994
# @@protoc_insertion_point(module_scope)
