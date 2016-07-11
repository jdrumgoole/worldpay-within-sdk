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
@Generated(value = "Autogenerated by Thrift Compiler (0.9.3)", date = "2016-07-11")
public class PricePerUnit implements org.apache.thrift.TBase<PricePerUnit, PricePerUnit._Fields>, java.io.Serializable, Cloneable, Comparable<PricePerUnit> {
  private static final org.apache.thrift.protocol.TStruct STRUCT_DESC = new org.apache.thrift.protocol.TStruct("PricePerUnit");

  private static final org.apache.thrift.protocol.TField AMOUNT_FIELD_DESC = new org.apache.thrift.protocol.TField("amount", org.apache.thrift.protocol.TType.I32, (short)1);
  private static final org.apache.thrift.protocol.TField CURRENCY_CODE_FIELD_DESC = new org.apache.thrift.protocol.TField("currencyCode", org.apache.thrift.protocol.TType.STRING, (short)2);

  private static final Map<Class<? extends IScheme>, SchemeFactory> schemes = new HashMap<Class<? extends IScheme>, SchemeFactory>();
  static {
    schemes.put(StandardScheme.class, new PricePerUnitStandardSchemeFactory());
    schemes.put(TupleScheme.class, new PricePerUnitTupleSchemeFactory());
  }

  public int amount; // required
  public String currencyCode; // required

  /** The set of fields this struct contains, along with convenience methods for finding and manipulating them. */
  public enum _Fields implements org.apache.thrift.TFieldIdEnum {
    AMOUNT((short)1, "amount"),
    CURRENCY_CODE((short)2, "currencyCode");

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
        case 1: // AMOUNT
          return AMOUNT;
        case 2: // CURRENCY_CODE
          return CURRENCY_CODE;
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
  private static final int __AMOUNT_ISSET_ID = 0;
  private byte __isset_bitfield = 0;
  public static final Map<_Fields, org.apache.thrift.meta_data.FieldMetaData> metaDataMap;
  static {
    Map<_Fields, org.apache.thrift.meta_data.FieldMetaData> tmpMap = new EnumMap<_Fields, org.apache.thrift.meta_data.FieldMetaData>(_Fields.class);
    tmpMap.put(_Fields.AMOUNT, new org.apache.thrift.meta_data.FieldMetaData("amount", org.apache.thrift.TFieldRequirementType.DEFAULT, 
        new org.apache.thrift.meta_data.FieldValueMetaData(org.apache.thrift.protocol.TType.I32)));
    tmpMap.put(_Fields.CURRENCY_CODE, new org.apache.thrift.meta_data.FieldMetaData("currencyCode", org.apache.thrift.TFieldRequirementType.DEFAULT, 
        new org.apache.thrift.meta_data.FieldValueMetaData(org.apache.thrift.protocol.TType.STRING)));
    metaDataMap = Collections.unmodifiableMap(tmpMap);
    org.apache.thrift.meta_data.FieldMetaData.addStructMetaDataMap(PricePerUnit.class, metaDataMap);
  }

  public PricePerUnit() {
  }

  public PricePerUnit(
    int amount,
    String currencyCode)
  {
    this();
    this.amount = amount;
    setAmountIsSet(true);
    this.currencyCode = currencyCode;
  }

  /**
   * Performs a deep copy on <i>other</i>.
   */
  public PricePerUnit(PricePerUnit other) {
    __isset_bitfield = other.__isset_bitfield;
    this.amount = other.amount;
    if (other.isSetCurrencyCode()) {
      this.currencyCode = other.currencyCode;
    }
  }

  public PricePerUnit deepCopy() {
    return new PricePerUnit(this);
  }

  @Override
  public void clear() {
    setAmountIsSet(false);
    this.amount = 0;
    this.currencyCode = null;
  }

  public int getAmount() {
    return this.amount;
  }

  public PricePerUnit setAmount(int amount) {
    this.amount = amount;
    setAmountIsSet(true);
    return this;
  }

  public void unsetAmount() {
    __isset_bitfield = EncodingUtils.clearBit(__isset_bitfield, __AMOUNT_ISSET_ID);
  }

  /** Returns true if field amount is set (has been assigned a value) and false otherwise */
  public boolean isSetAmount() {
    return EncodingUtils.testBit(__isset_bitfield, __AMOUNT_ISSET_ID);
  }

  public void setAmountIsSet(boolean value) {
    __isset_bitfield = EncodingUtils.setBit(__isset_bitfield, __AMOUNT_ISSET_ID, value);
  }

  public String getCurrencyCode() {
    return this.currencyCode;
  }

  public PricePerUnit setCurrencyCode(String currencyCode) {
    this.currencyCode = currencyCode;
    return this;
  }

  public void unsetCurrencyCode() {
    this.currencyCode = null;
  }

  /** Returns true if field currencyCode is set (has been assigned a value) and false otherwise */
  public boolean isSetCurrencyCode() {
    return this.currencyCode != null;
  }

  public void setCurrencyCodeIsSet(boolean value) {
    if (!value) {
      this.currencyCode = null;
    }
  }

  public void setFieldValue(_Fields field, Object value) {
    switch (field) {
    case AMOUNT:
      if (value == null) {
        unsetAmount();
      } else {
        setAmount((Integer)value);
      }
      break;

    case CURRENCY_CODE:
      if (value == null) {
        unsetCurrencyCode();
      } else {
        setCurrencyCode((String)value);
      }
      break;

    }
  }

  public Object getFieldValue(_Fields field) {
    switch (field) {
    case AMOUNT:
      return getAmount();

    case CURRENCY_CODE:
      return getCurrencyCode();

    }
    throw new IllegalStateException();
  }

  /** Returns true if field corresponding to fieldID is set (has been assigned a value) and false otherwise */
  public boolean isSet(_Fields field) {
    if (field == null) {
      throw new IllegalArgumentException();
    }

    switch (field) {
    case AMOUNT:
      return isSetAmount();
    case CURRENCY_CODE:
      return isSetCurrencyCode();
    }
    throw new IllegalStateException();
  }

  @Override
  public boolean equals(Object that) {
    if (that == null)
      return false;
    if (that instanceof PricePerUnit)
      return this.equals((PricePerUnit)that);
    return false;
  }

  public boolean equals(PricePerUnit that) {
    if (that == null)
      return false;

    boolean this_present_amount = true;
    boolean that_present_amount = true;
    if (this_present_amount || that_present_amount) {
      if (!(this_present_amount && that_present_amount))
        return false;
      if (this.amount != that.amount)
        return false;
    }

    boolean this_present_currencyCode = true && this.isSetCurrencyCode();
    boolean that_present_currencyCode = true && that.isSetCurrencyCode();
    if (this_present_currencyCode || that_present_currencyCode) {
      if (!(this_present_currencyCode && that_present_currencyCode))
        return false;
      if (!this.currencyCode.equals(that.currencyCode))
        return false;
    }

    return true;
  }

  @Override
  public int hashCode() {
    List<Object> list = new ArrayList<Object>();

    boolean present_amount = true;
    list.add(present_amount);
    if (present_amount)
      list.add(amount);

    boolean present_currencyCode = true && (isSetCurrencyCode());
    list.add(present_currencyCode);
    if (present_currencyCode)
      list.add(currencyCode);

    return list.hashCode();
  }

  @Override
  public int compareTo(PricePerUnit other) {
    if (!getClass().equals(other.getClass())) {
      return getClass().getName().compareTo(other.getClass().getName());
    }

    int lastComparison = 0;

    lastComparison = Boolean.valueOf(isSetAmount()).compareTo(other.isSetAmount());
    if (lastComparison != 0) {
      return lastComparison;
    }
    if (isSetAmount()) {
      lastComparison = org.apache.thrift.TBaseHelper.compareTo(this.amount, other.amount);
      if (lastComparison != 0) {
        return lastComparison;
      }
    }
    lastComparison = Boolean.valueOf(isSetCurrencyCode()).compareTo(other.isSetCurrencyCode());
    if (lastComparison != 0) {
      return lastComparison;
    }
    if (isSetCurrencyCode()) {
      lastComparison = org.apache.thrift.TBaseHelper.compareTo(this.currencyCode, other.currencyCode);
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
    StringBuilder sb = new StringBuilder("PricePerUnit(");
    boolean first = true;

    sb.append("amount:");
    sb.append(this.amount);
    first = false;
    if (!first) sb.append(", ");
    sb.append("currencyCode:");
    if (this.currencyCode == null) {
      sb.append("null");
    } else {
      sb.append(this.currencyCode);
    }
    first = false;
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

  private static class PricePerUnitStandardSchemeFactory implements SchemeFactory {
    public PricePerUnitStandardScheme getScheme() {
      return new PricePerUnitStandardScheme();
    }
  }

  private static class PricePerUnitStandardScheme extends StandardScheme<PricePerUnit> {

    public void read(org.apache.thrift.protocol.TProtocol iprot, PricePerUnit struct) throws org.apache.thrift.TException {
      org.apache.thrift.protocol.TField schemeField;
      iprot.readStructBegin();
      while (true)
      {
        schemeField = iprot.readFieldBegin();
        if (schemeField.type == org.apache.thrift.protocol.TType.STOP) { 
          break;
        }
        switch (schemeField.id) {
          case 1: // AMOUNT
            if (schemeField.type == org.apache.thrift.protocol.TType.I32) {
              struct.amount = iprot.readI32();
              struct.setAmountIsSet(true);
            } else { 
              org.apache.thrift.protocol.TProtocolUtil.skip(iprot, schemeField.type);
            }
            break;
          case 2: // CURRENCY_CODE
            if (schemeField.type == org.apache.thrift.protocol.TType.STRING) {
              struct.currencyCode = iprot.readString();
              struct.setCurrencyCodeIsSet(true);
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

    public void write(org.apache.thrift.protocol.TProtocol oprot, PricePerUnit struct) throws org.apache.thrift.TException {
      struct.validate();

      oprot.writeStructBegin(STRUCT_DESC);
      oprot.writeFieldBegin(AMOUNT_FIELD_DESC);
      oprot.writeI32(struct.amount);
      oprot.writeFieldEnd();
      if (struct.currencyCode != null) {
        oprot.writeFieldBegin(CURRENCY_CODE_FIELD_DESC);
        oprot.writeString(struct.currencyCode);
        oprot.writeFieldEnd();
      }
      oprot.writeFieldStop();
      oprot.writeStructEnd();
    }

  }

  private static class PricePerUnitTupleSchemeFactory implements SchemeFactory {
    public PricePerUnitTupleScheme getScheme() {
      return new PricePerUnitTupleScheme();
    }
  }

  private static class PricePerUnitTupleScheme extends TupleScheme<PricePerUnit> {

    @Override
    public void write(org.apache.thrift.protocol.TProtocol prot, PricePerUnit struct) throws org.apache.thrift.TException {
      TTupleProtocol oprot = (TTupleProtocol) prot;
      BitSet optionals = new BitSet();
      if (struct.isSetAmount()) {
        optionals.set(0);
      }
      if (struct.isSetCurrencyCode()) {
        optionals.set(1);
      }
      oprot.writeBitSet(optionals, 2);
      if (struct.isSetAmount()) {
        oprot.writeI32(struct.amount);
      }
      if (struct.isSetCurrencyCode()) {
        oprot.writeString(struct.currencyCode);
      }
    }

    @Override
    public void read(org.apache.thrift.protocol.TProtocol prot, PricePerUnit struct) throws org.apache.thrift.TException {
      TTupleProtocol iprot = (TTupleProtocol) prot;
      BitSet incoming = iprot.readBitSet(2);
      if (incoming.get(0)) {
        struct.amount = iprot.readI32();
        struct.setAmountIsSet(true);
      }
      if (incoming.get(1)) {
        struct.currencyCode = iprot.readString();
        struct.setCurrencyCodeIsSet(true);
      }
    }
  }

}

