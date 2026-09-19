class $classObject.getName() < ApplicationRecord
#set( $includePrimaryKeys = false )
#set( $includeTypes = false )
#set( $includeAssociations = false )
#set( $delim = ", :" )
#set( $suffix = "" )
#foreach( $enumAsAttribute in $classObject.getEnumerators() )
#set( $enum = $aib.getEnumClassObject( $enumAsAttribute.getType() ) )
#set( $attributes = $enum.getAttributesAsString( $includePrimaryKeys, $includeTypes, $includeAssociations, $delim, $suffix ) )
  enum $enumAsAttribute.getName(): [:${attributes}]
#end##foreach( $enum in $classObject.getEnumerators() )

#foreach( $attribute in $classObject.getAttributesOnly(false,false) )
#if ( $aib.isValueObject( $attribute.getType() ) )
#set( $valueObject = $aib.getValueObject( $attribute.getType()  ) )
#set( $className = $valueObject.getName() )
#set( $lowercaseClassName = ${Utils.lowercaseFirstLetter(${className})} )

  composed_of :${lowercaseClassName},
    class_name: "${}className}",
    mapping: [
#set( $voAttributes = $valueObject.getAttributesOnly(false,false) )
#set( $voAttributesSize = voAttributes.size() )
#foreach( $attrib in $voAttributes )
#set( $attributeName = ${Utils.lowercaseFirstLetter(${attrib.getName()})} )
#set( $mapping = "%w[${lowercaseClassName}_${attributeName} ${attributeName}]")
#if ($velocityCount < $voAttributesSize )
#set( $createDecl = "${$mapping}, " )
#end
      ${mapping}
#end
    ]
#end
#end##foreach( $valueObject in $classObject.getValueObjects() )
#declareAssociations( $classObject )
end
