import React, { Component } from 'react'
import ItemService from '../services/ItemService'

class ViewItemComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            item: {}
        }
    }

    componentDidMount(){
        ItemService.getItemById(this.state.id).then( res => {
            this.setState({item: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View Item Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> itemNumber:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.item.itemNumber }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> name:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.item.name }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> standardCost:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.item.standardCost }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> weight:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.item.weight }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> asSerialControlled:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.item.asSerialControlled }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> ItemType:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.item.itemType }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> ProcurementType:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.item.procurementType }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> UnitOfMeasure:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.item.unitOfMeasure }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> LifecycleStatus:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.item.lifecycleStatus }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewItemComponent
