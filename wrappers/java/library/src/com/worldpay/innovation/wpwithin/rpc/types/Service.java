/**
 * Autogenerated by Thrift Compiler (0.9.3)
 *
 * DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING
 *  @generated
 */
package com.worldpay.innovation.wpwithin.rpc.types;

import org.apache.thrift.scheme.IScheme;
import org.apache.thrift.scheme.SchemeFactory;
import org.apache.thrift.scheme.StandardScheme;

import org.apache.thrift.scheme.TupleScheme;
import org.apache.thrift.protocol.TTupleProtocol;
import org.apache.thrift.protocol.TProtocolException;
import org.apache.thrift.EncodingUtils;
import org.apache.thrift.TException;
import org.apache.thrift.async.AsyncMethodCallback;
import org.apache.thrift.server.AbstractNonblockingServer.*;
import java.util.List;
import java.util.ArrayList;
import java.util.Map;
import java.util.HashMap;
import java.util.EnumMap;
import java.util.Set;
import java.util.HashSet;
import java.util.EnumSet;
import java.util.Collections;
import java.util.BitSet;
import java.nio.ByteBuffer;
import java.util.Arrays;
import javax.annotation.Generated;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

@SuppressWarnings({"cast", "rawtypes", "serial", "unchecked"})
@Generated(value = "Autogenerated by Thrift Compiler (0.9.3)", date = "2016-09-03")
public class Service implements org.apache.thrift.TBase<Service, Service._Fields>, java.io.Serializable, Cloneable, Comparable<Service> {
  private static final org.apache.thrift.protocol.TStruct STRUCT_DESC = new org.apache.thrift.protocol.TStruct("Service");

  private static final org.apache.thrift.protocol.TField ID_FIELD_DESC = new org.apache.thrift.protocol.TField("id", org.apache.thrift.protocol.TType.I32, (short)1);
  private static final org.apache.thrift.protocol.TField NAME_FIELD_DESC = new org.apache.thrift.protocol.TField("name", org.apache.thrift.protocol.TType.STRING, (short)2);
  private static final org.apache.thrift.protocol.TField DESCRIPTION_FIELD_DESC = new org.apache.thrift.protocol.TField("description", org.apache.thrift.protocol.TType.STRING, (short)3);
  private static final org.apache.thrift.protocol.TField PRICES_FIELD_DESC = new org.apache.thrift.protocol.TField("prices", org.apache.thrift.protocol.TType.MAP, (short)4);

  private static final Map<Class<? extends IScheme>, SchemeFactory> schemes = new HashMap<Class<? extends IScheme>, SchemeFactory>();
  static {
    schemes.put(StandardScheme.class, new ServiceStandardSchemeFactory());
    schemes.put(TupleScheme.class, new ServiceTupleSchemeFactory());
  }

  public int id; // required
  public String name; // required
  public String description; // required
  public Map<Integer,Price> prices; // optional

  /** The set of fields this struct contains, along with convenience methods for finding and manipulating them. */
  public enum _Fields implements org.apache.thrift.TFieldIdEnum {
    ID((short)1, "id"),
    NAME((short)2, "name"),
    DESCRIPTION((short)3, "description"),
    PRICES((short)4, "prices");

    private static final Map<String, _Fields> byName = new HashMap<String, _Fields>();

    static {
      for (_Fields field : EnumSet.allOf(_Fields.class)) {
        byName.put(field.getFieldName(), field);
      }
    }

    /**
     * Find the _Fields constant that matches fieldId, or null if its not found.
     */
    public static _Fields findByThriftId(int fieldId) {
      switch(fieldId) {
        case 1: // ID
          return ID;
        case 2: // NAME
          return NAME;
        case 3: // DESCRIPTION
          return DESCRIPTION;
        case 4: // PRICES
          return PRICES;
        default:
          return null;
      }
    }

    /**
     * Find the _Fields constant that matches fieldId, throwing an exception
     * if it is not found.
     */
    public static _Fields findByThriftIdOrThrow(int fieldId) {
      _Fields fields = findByThriftId(fieldId);
      if (fields == null) throw new IllegalArgumentException("Field " + fieldId + " doesn't exist!");
      return fields;
    }

    /**
     * Find the _Fields constant that matches name, or null if its not found.
     */
    public static _Fields findByName(String name) {
      return byName.get(name);
    }

    private final short _thriftId;
    private final String _fieldName;

    _Fields(short thriftId, String fieldName) {
      _thriftId = thriftId;
      _fieldName = fieldName;
    }

    public short getThriftFieldId() {
      return _thriftId;
    }

    public String getFieldName() {
      return _fieldName;
    }
  }

  // isset id assignments
  private static final int __ID_ISSET_ID = 0;
  private byte __isset_bitfield = 0;
  private static final _Fields optionals[] = {_Fields.PRICES};
  public static final Map<_Fields, org.apache.thrift.meta_data.FieldMetaData> metaDataMap;
  static {
    Map<_Fields, org.apache.thrift.meta_data.FieldMetaData> tmpMap = new EnumMap<_Fields, org.apache.thrift.meta_data.FieldMetaData>(_Fields.class);
    tmpMap.put(_Fields.ID, new org.apache.thrift.meta_data.FieldMetaData("id", org.apache.thrift.TFieldRequirementType.DEFAULT, 
        new org.apache.thrift.meta_data.FieldValueMetaData(org.apache.thrift.protocol.TType.I32)));
    tmpMap.put(_Fields.NAME, new org.apache.thrift.meta_data.FieldMetaData("name", org.apache.thrift.TFieldRequirementType.DEFAULT, 
        new org.apache.thrift.meta_data.FieldValueMetaData(org.apache.thrift.protocol.TType.STRING)));
    tmpMap.put(_Fields.DESCRIPTION, new org.apache.thrift.meta_data.FieldMetaData("description", org.apache.thrift.TFieldRequirementType.DEFAULT, 
        new org.apache.thrift.meta_data.FieldValueMetaData(org.apache.thrift.protocol.TType.STRING)));
    tmpMap.put(_Fields.PRICES, new org.apache.thrift.meta_data.FieldMetaData("prices", org.apache.thrift.TFieldRequirementType.OPTIONAL, 
        new org.apache.thrift.meta_data.MapMetaData(org.apache.thrift.protocol.TType.MAP, 
            new org.apache.thrift.meta_data.FieldValueMetaData(org.apache.thrift.protocol.TType.I32), 
            new org.apache.thrift.meta_data.FieldValueMetaData(org.apache.thrift.protocol.TType.STRUCT            , "Price"))));
    metaDataMap = Collections.unmodifiableMap(tmpMap);
    org.apache.thrift.meta_data.FieldMetaData.addStructMetaDataMap(Service.class, metaDataMap);
  }

  public Service() {
  }

  public Service(
    int id,
    String name,
    String description)
  {
    this();
    this.id = id;
    setIdIsSet(true);
    this.name = name;
    this.description = description;
  }

  /**
   * Performs a deep copy on <i>other</i>.
   */
  public Service(Service other) {
    __isset_bitfield = other.__isset_bitfield;
    this.id = other.id;
    if (other.isSetName()) {
      this.name = other.name;
    }
    if (other.isSetDescription()) {
      this.description = other.description;
    }
    if (other.isSetPrices()) {
      Map<Integer,Price> __this__prices = new HashMap<Integer,Price>(other.prices.size());
      for (Map.Entry<Integer, Price> other_element : other.prices.entrySet()) {

        Integer other_element_key = other_element.getKey();
        Price other_element_value = other_element.getValue();

        Integer __this__prices_copy_key = other_element_key;

        Price __this__prices_copy_value = other_element_value;

        __this__prices.put(__this__prices_copy_key, __this__prices_copy_value);
      }
      this.prices = __this__prices;
    }
  }

  public Service deepCopy() {
    return new Service(this);
  }

  @Override
  public void clear() {
    setIdIsSet(false);
    this.id = 0;
    this.name = null;
    this.description = null;
    this.prices = null;
  }

  public int getId() {
    return this.id;
  }

  public Service setId(int id) {
    this.id = id;
    setIdIsSet(true);
    return this;
  }

  public void unsetId() {
    __isset_bitfield = EncodingUtils.clearBit(__isset_bitfield, __ID_ISSET_ID);
  }

  /** Returns true if field id is set (has been assigned a value) and false otherwise */
  public boolean isSetId() {
    return EncodingUtils.testBit(__isset_bitfield, __ID_ISSET_ID);
  }

  public void setIdIsSet(boolean value) {
    __isset_bitfield = EncodingUtils.setBit(__isset_bitfield, __ID_ISSET_ID, value);
  }

  public String getName() {
    return this.name;
  }

  public Service setName(String name) {
    this.name = name;
    return this;
  }

  public void unsetName() {
    this.name = null;
  }

  /** Returns true if field name is set (has been assigned a value) and false otherwise */
  public boolean isSetName() {
    return this.name != null;
  }

  public void setNameIsSet(boolean value) {
    if (!value) {
      this.name = null;
    }
  }

  public String getDescription() {
    return this.description;
  }

  public Service setDescription(String description) {
    this.description = description;
    return this;
  }

  public void unsetDescription() {
    this.description = null;
  }

  /** Returns true if field description is set (has been assigned a value) and false otherwise */
  public boolean isSetDescription() {
    return this.description != null;
  }

  public void setDescriptionIsSet(boolean value) {
    if (!value) {
      this.description = null;
    }
  }

  public int getPricesSize() {
    return (this.prices == null) ? 0 : this.prices.size();
  }

  public void putToPrices(int key, Price val) {
    if (this.prices == null) {
      this.prices = new HashMap<Integer,Price>();
    }
    this.prices.put(key, val);
  }

  public Map<Integer,Price> getPrices() {
    return this.prices;
  }

  public Service setPrices(Map<Integer,Price> prices) {
    this.prices = prices;
    return this;
  }

  public void unsetPrices() {
    this.prices = null;
  }

  /** Returns true if field prices is set (has been assigned a value) and false otherwise */
  public boolean isSetPrices() {
    return this.prices != null;
  }

  public void setPricesIsSet(boolean value) {
    if (!value) {
      this.prices = null;
    }
  }

  public void setFieldValue(_Fields field, Object value) {
    switch (field) {
    case ID:
      if (value == null) {
        unsetId();
      } else {
        setId((Integer)value);
      }
      break;

    case NAME:
      if (value == null) {
        unsetName();
      } else {
        setName((String)value);
      }
      break;

    case DESCRIPTION:
      if (value == null) {
        unsetDescription();
      } else {
        setDescription((String)value);
      }
      break;

    case PRICES:
      if (value == null) {
        unsetPrices();
      } else {
        setPrices((Map<Integer,Price>)value);
      }
      break;

    }
  }

  public Object getFieldValue(_Fields field) {
    switch (field) {
    case ID:
      return getId();

    case NAME:
      return getName();

    case DESCRIPTION:
      return getDescription();

    case PRICES:
      return getPrices();

    }
    throw new IllegalStateException();
  }

  /** Returns true if field corresponding to fieldID is set (has been assigned a value) and false otherwise */
  public boolean isSet(_Fields field) {
    if (field == null) {
      throw new IllegalArgumentException();
    }

    switch (field) {
    case ID:
      return isSetId();
    case NAME:
      return isSetName();
    case DESCRIPTION:
      return isSetDescription();
    case PRICES:
      return isSetPrices();
    }
    throw new IllegalStateException();
  }

  @Override
  public boolean equals(Object that) {
    if (that == null)
      return false;
    if (that instanceof Service)
      return this.equals((Service)that);
    return false;
  }

  public boolean equals(Service that) {
    if (that == null)
      return false;

    boolean this_present_id = true;
    boolean that_present_id = true;
    if (this_present_id || that_present_id) {
      if (!(this_present_id && that_present_id))
        return false;
      if (this.id != that.id)
        return false;
    }

    boolean this_present_name = true && this.isSetName();
    boolean that_present_name = true && that.isSetName();
    if (this_present_name || that_present_name) {
      if (!(this_present_name && that_present_name))
        return false;
      if (!this.name.equals(that.name))
        return false;
    }

    boolean this_present_description = true && this.isSetDescription();
    boolean that_present_description = true && that.isSetDescription();
    if (this_present_description || that_present_description) {
      if (!(this_present_description && that_present_description))
        return false;
      if (!this.description.equals(that.description))
        return false;
    }

    boolean this_present_prices = true && this.isSetPrices();
    boolean that_present_prices = true && that.isSetPrices();
    if (this_present_prices || that_present_prices) {
      if (!(this_present_prices && that_present_prices))
        return false;
      if (!this.prices.equals(that.prices))
        return false;
    }

    return true;
  }

  @Override
  public int hashCode() {
    List<Object> list = new ArrayList<Object>();

    boolean present_id = true;
    list.add(present_id);
    if (present_id)
      list.add(id);

    boolean present_name = true && (isSetName());
    list.add(present_name);
    if (present_name)
      list.add(name);

    boolean present_description = true && (isSetDescription());
    list.add(present_description);
    if (present_description)
      list.add(description);

    boolean present_prices = true && (isSetPrices());
    list.add(present_prices);
    if (present_prices)
      list.add(prices);

    return list.hashCode();
  }

  @Override
  public int compareTo(Service other) {
    if (!getClass().equals(other.getClass())) {
      return getClass().getName().compareTo(other.getClass().getName());
    }

    int lastComparison = 0;

    lastComparison = Boolean.valueOf(isSetId()).compareTo(other.isSetId());
    if (lastComparison != 0) {
      return lastComparison;
    }
    if (isSetId()) {
      lastComparison = org.apache.thrift.TBaseHelper.compareTo(this.id, other.id);
      if (lastComparison != 0) {
        return lastComparison;
      }
    }
    lastComparison = Boolean.valueOf(isSetName()).compareTo(other.isSetName());
    if (lastComparison != 0) {
      return lastComparison;
    }
    if (isSetName()) {
      lastComparison = org.apache.thrift.TBaseHelper.compareTo(this.name, other.name);
      if (lastComparison != 0) {
        return lastComparison;
      }
    }
    lastComparison = Boolean.valueOf(isSetDescription()).compareTo(other.isSetDescription());
    if (lastComparison != 0) {
      return lastComparison;
    }
    if (isSetDescription()) {
      lastComparison = org.apache.thrift.TBaseHelper.compareTo(this.description, other.description);
      if (lastComparison != 0) {
        return lastComparison;
      }
    }
    lastComparison = Boolean.valueOf(isSetPrices()).compareTo(other.isSetPrices());
    if (lastComparison != 0) {
      return lastComparison;
    }
    if (isSetPrices()) {
      lastComparison = org.apache.thrift.TBaseHelper.compareTo(this.prices, other.prices);
      if (lastComparison != 0) {
        return lastComparison;
      }
    }
    return 0;
  }

  public _Fields fieldForId(int fieldId) {
    return _Fields.findByThriftId(fieldId);
  }

  public void read(org.apache.thrift.protocol.TProtocol iprot) throws org.apache.thrift.TException {
    schemes.get(iprot.getScheme()).getScheme().read(iprot, this);
  }

  public void write(org.apache.thrift.protocol.TProtocol oprot) throws org.apache.thrift.TException {
    schemes.get(oprot.getScheme()).getScheme().write(oprot, this);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder("Service(");
    boolean first = true;

    sb.append("id:");
    sb.append(this.id);
    first = false;
    if (!first) sb.append(", ");
    sb.append("name:");
    if (this.name == null) {
      sb.append("null");
    } else {
      sb.append(this.name);
    }
    first = false;
    if (!first) sb.append(", ");
    sb.append("description:");
    if (this.description == null) {
      sb.append("null");
    } else {
      sb.append(this.description);
    }
    first = false;
    if (isSetPrices()) {
      if (!first) sb.append(", ");
      sb.append("prices:");
      if (this.prices == null) {
        sb.append("null");
      } else {
        sb.append(this.prices);
      }
      first = false;
    }
    sb.append(")");
    return sb.toString();
  }

  public void validate() throws org.apache.thrift.TException {
    // check for required fields
    // check for sub-struct validity
  }

  private void writeObject(java.io.ObjectOutputStream out) throws java.io.IOException {
    try {
      write(new org.apache.thrift.protocol.TCompactProtocol(new org.apache.thrift.transport.TIOStreamTransport(out)));
    } catch (org.apache.thrift.TException te) {
      throw new java.io.IOException(te);
    }
  }

  private void readObject(java.io.ObjectInputStream in) throws java.io.IOException, ClassNotFoundException {
    try {
      // it doesn't seem like you should have to do this, but java serialization is wacky, and doesn't call the default constructor.
      __isset_bitfield = 0;
      read(new org.apache.thrift.protocol.TCompactProtocol(new org.apache.thrift.transport.TIOStreamTransport(in)));
    } catch (org.apache.thrift.TException te) {
      throw new java.io.IOException(te);
    }
  }

  private static class ServiceStandardSchemeFactory implements SchemeFactory {
    public ServiceStandardScheme getScheme() {
      return new ServiceStandardScheme();
    }
  }

  private static class ServiceStandardScheme extends StandardScheme<Service> {

    public void read(org.apache.thrift.protocol.TProtocol iprot, Service struct) throws org.apache.thrift.TException {
      org.apache.thrift.protocol.TField schemeField;
      iprot.readStructBegin();
      while (true)
      {
        schemeField = iprot.readFieldBegin();
        if (schemeField.type == org.apache.thrift.protocol.TType.STOP) { 
          break;
        }
        switch (schemeField.id) {
          case 1: // ID
            if (schemeField.type == org.apache.thrift.protocol.TType.I32) {
              struct.id = iprot.readI32();
              struct.setIdIsSet(true);
            } else { 
              org.apache.thrift.protocol.TProtocolUtil.skip(iprot, schemeField.type);
            }
            break;
          case 2: // NAME
            if (schemeField.type == org.apache.thrift.protocol.TType.STRING) {
              struct.name = iprot.readString();
              struct.setNameIsSet(true);
            } else { 
              org.apache.thrift.protocol.TProtocolUtil.skip(iprot, schemeField.type);
            }
            break;
          case 3: // DESCRIPTION
            if (schemeField.type == org.apache.thrift.protocol.TType.STRING) {
              struct.description = iprot.readString();
              struct.setDescriptionIsSet(true);
            } else { 
              org.apache.thrift.protocol.TProtocolUtil.skip(iprot, schemeField.type);
            }
            break;
          case 4: // PRICES
            if (schemeField.type == org.apache.thrift.protocol.TType.MAP) {
              {
                org.apache.thrift.protocol.TMap _map0 = iprot.readMapBegin();
                struct.prices = new HashMap<Integer,Price>(2*_map0.size);
                int _key1;
                Price _val2;
                for (int _i3 = 0; _i3 < _map0.size; ++_i3)
                {
                  _key1 = iprot.readI32();
                  _val2 = new Price();
                  _val2.read(iprot);
                  struct.prices.put(_key1, _val2);
                }
                iprot.readMapEnd();
              }
              struct.setPricesIsSet(true);
            } else { 
              org.apache.thrift.protocol.TProtocolUtil.skip(iprot, schemeField.type);
            }
            break;
          default:
            org.apache.thrift.protocol.TProtocolUtil.skip(iprot, schemeField.type);
        }
        iprot.readFieldEnd();
      }
      iprot.readStructEnd();

      // check for required fields of primitive type, which can't be checked in the validate method
      struct.validate();
    }

    public void write(org.apache.thrift.protocol.TProtocol oprot, Service struct) throws org.apache.thrift.TException {
      struct.validate();

      oprot.writeStructBegin(STRUCT_DESC);
      oprot.writeFieldBegin(ID_FIELD_DESC);
      oprot.writeI32(struct.id);
      oprot.writeFieldEnd();
      if (struct.name != null) {
        oprot.writeFieldBegin(NAME_FIELD_DESC);
        oprot.writeString(struct.name);
        oprot.writeFieldEnd();
      }
      if (struct.description != null) {
        oprot.writeFieldBegin(DESCRIPTION_FIELD_DESC);
        oprot.writeString(struct.description);
        oprot.writeFieldEnd();
      }
      if (struct.prices != null) {
        if (struct.isSetPrices()) {
          oprot.writeFieldBegin(PRICES_FIELD_DESC);
          {
            oprot.writeMapBegin(new org.apache.thrift.protocol.TMap(org.apache.thrift.protocol.TType.I32, org.apache.thrift.protocol.TType.STRUCT, struct.prices.size()));
            for (Map.Entry<Integer, Price> _iter4 : struct.prices.entrySet())
            {
              oprot.writeI32(_iter4.getKey());
              _iter4.getValue().write(oprot);
            }
            oprot.writeMapEnd();
          }
          oprot.writeFieldEnd();
        }
      }
      oprot.writeFieldStop();
      oprot.writeStructEnd();
    }

  }

  private static class ServiceTupleSchemeFactory implements SchemeFactory {
    public ServiceTupleScheme getScheme() {
      return new ServiceTupleScheme();
    }
  }

  private static class ServiceTupleScheme extends TupleScheme<Service> {

    @Override
    public void write(org.apache.thrift.protocol.TProtocol prot, Service struct) throws org.apache.thrift.TException {
      TTupleProtocol oprot = (TTupleProtocol) prot;
      BitSet optionals = new BitSet();
      if (struct.isSetId()) {
        optionals.set(0);
      }
      if (struct.isSetName()) {
        optionals.set(1);
      }
      if (struct.isSetDescription()) {
        optionals.set(2);
      }
      if (struct.isSetPrices()) {
        optionals.set(3);
      }
      oprot.writeBitSet(optionals, 4);
      if (struct.isSetId()) {
        oprot.writeI32(struct.id);
      }
      if (struct.isSetName()) {
        oprot.writeString(struct.name);
      }
      if (struct.isSetDescription()) {
        oprot.writeString(struct.description);
      }
      if (struct.isSetPrices()) {
        {
          oprot.writeI32(struct.prices.size());
          for (Map.Entry<Integer, Price> _iter5 : struct.prices.entrySet())
          {
            oprot.writeI32(_iter5.getKey());
            _iter5.getValue().write(oprot);
          }
        }
      }
    }

    @Override
    public void read(org.apache.thrift.protocol.TProtocol prot, Service struct) throws org.apache.thrift.TException {
      TTupleProtocol iprot = (TTupleProtocol) prot;
      BitSet incoming = iprot.readBitSet(4);
      if (incoming.get(0)) {
        struct.id = iprot.readI32();
        struct.setIdIsSet(true);
      }
      if (incoming.get(1)) {
        struct.name = iprot.readString();
        struct.setNameIsSet(true);
      }
      if (incoming.get(2)) {
        struct.description = iprot.readString();
        struct.setDescriptionIsSet(true);
      }
      if (incoming.get(3)) {
        {
          org.apache.thrift.protocol.TMap _map6 = new org.apache.thrift.protocol.TMap(org.apache.thrift.protocol.TType.I32, org.apache.thrift.protocol.TType.STRUCT, iprot.readI32());
          struct.prices = new HashMap<Integer,Price>(2*_map6.size);
          int _key7;
          Price _val8;
          for (int _i9 = 0; _i9 < _map6.size; ++_i9)
          {
            _key7 = iprot.readI32();
            _val8 = new Price();
            _val8.read(iprot);
            struct.prices.put(_key7, _val8);
          }
        }
        struct.setPricesIsSet(true);
      }
    }
  }

}

