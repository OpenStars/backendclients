rm -rf ./gen-cpp
#/data/data/zserver/thrift9.3/bin/thrift -gen cpp -r ./bigsetgenericdata.thrift
#/data/data/zserver/thrift9.3/bin/thrift -gen cpp -r ./idgen.thrift
name=bigsetgenericdata
#sed -i '/virtual ~/d' gen-cpp/"$name"_types.h
#sed -i '/::~/,3d' gen-cpp/"$name"_types.cpp
#sed -i  's/virtual ~/~/' gen-cpp/"$name"_types.h

../../../../contribs/ApacheThrift/bin/thrift -gen cpp bigsetgenericdata.thrift
../../../../contribs/ApacheThrift/bin/thrift -gen cpp idgen.thrift