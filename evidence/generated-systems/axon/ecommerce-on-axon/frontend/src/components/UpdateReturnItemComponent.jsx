import React, { Component } from 'react'
import ReturnItemService from '../services/ReturnItemService';

class UpdateReturnItemComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                quantity: '',
                reason: '',
                condition: ''
        }
        this.updateReturnItem = this.updateReturnItem.bind(this);

        this.changequantityHandler = this.changequantityHandler.bind(this);
        this.changeReasonHandler = this.changeReasonHandler.bind(this);
        this.changeConditionHandler = this.changeConditionHandler.bind(this);
    }

    componentDidMount(){
        ReturnItemService.getReturnItemById(this.state.id).then( (res) =>{
            let returnItem = res.data;
            this.setState({
                quantity: returnItem.quantity,
                reason: returnItem.reason,
                condition: returnItem.condition
            });
        });
    }

    updateReturnItem = (e) => {
        e.preventDefault();
        let returnItem = {
            returnItemId: this.state.id,
            quantity: this.state.quantity,
            reason: this.state.reason,
            condition: this.state.condition
        };
        console.log('returnItem => ' + JSON.stringify(returnItem));
        console.log('id => ' + JSON.stringify(this.state.id));
        ReturnItemService.updateReturnItem(returnItem).then( res => {
            this.props.history.push('/returnItems');
        });
    }

    changequantityHandler= (event) => {
        this.setState({quantity: event.target.value});
    }
    changeReasonHandler= (event) => {
        this.setState({reason: event.target.value});
    }
    changeConditionHandler= (event) => {
        this.setState({condition: event.target.value});
    }

    cancel(){
        this.props.history.push('/returnItems');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update ReturnItem</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> quantity: </label>
                                                <input type="number" placeholder="quantity" name="quantity" className="form-control" value={this.state.quantity} onChange={this.changequantityHandler}/>

                                            <label> Reason: </label>
                                                <select value={this.state.reason} onChange={this.changeReasonHandler}>
                      <option name="Reason" className="form-control" >
                          Defective
                      </option>
                      <option name="Reason" className="form-control" >
                          Damaged
                      </option>
                      <option name="Reason" className="form-control" >
                          NotAsDescribed
                      </option>
                      <option name="Reason" className="form-control" >
                          WrongItem
                      </option>
                      <option name="Reason" className="form-control" >
                          NoLongerNeeded
                      </option>
                      <option name="Reason" className="form-control" >
                          SizeFitIssue
                      </option>
                      <option name="Reason" className="form-control" >
                          Other
                      </option>
                    </select>

                                            <label> Condition: </label>
                                                <select value={this.state.condition} onChange={this.changeConditionHandler}>
                      <option name="Condition" className="form-control" >
                          New
                      </option>
                      <option name="Condition" className="form-control" >
                          OpenBox
                      </option>
                      <option name="Condition" className="form-control" >
                          Used
                      </option>
                      <option name="Condition" className="form-control" >
                          Damaged
                      </option>
                      <option name="Condition" className="form-control" >
                          MissingParts
                      </option>
                    </select>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateReturnItem}>Save</button>
                                        <button className="btn btn-danger" onClick={this.cancel.bind(this)} style={{marginLeft: "10px"}}>Cancel</button>
                                    </form>
                                </div>
                            </div>
                        </div>

                   </div>
            </div>
        )
    }
}

export default UpdateReturnItemComponent
