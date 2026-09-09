import React, { Component } from 'react'
import ReturnItemService from '../services/ReturnItemService';

class CreateReturnItemComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                quantity: '',
                reason: '',
                condition: ''
        }
        this.changequantityHandler = this.changequantityHandler.bind(this);
        this.changeReasonHandler = this.changeReasonHandler.bind(this);
        this.changeConditionHandler = this.changeConditionHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            ReturnItemService.getReturnItemById(this.state.id).then( (res) =>{
                let returnItem = res.data;
                this.setState({
                    quantity: returnItem.quantity,
                    reason: returnItem.reason,
                    condition: returnItem.condition
                });
            });
        }        
    }
    saveOrUpdateReturnItem = (e) => {
        e.preventDefault();
        let returnItem = {
                returnItemId: this.state.id,
                quantity: this.state.quantity,
                reason: this.state.reason,
                condition: this.state.condition
            };
        console.log('returnItem => ' + JSON.stringify(returnItem));

        // step 5
        if(this.state.id === '_add'){
            returnItem.returnItemId=''
            ReturnItemService.createReturnItem(returnItem).then(res =>{
                this.props.history.push('/returnItems');
            });
        }else{
            ReturnItemService.updateReturnItem(returnItem).then( res => {
                this.props.history.push('/returnItems');
            });
        }
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

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add ReturnItem</h3>
        }else{
            return <h3 className="text-center">Update ReturnItem</h3>
        }
    }
    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                {
                                    this.getTitle()
                                }
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> quantity:&emsp; </label>
                                                <input type="number" placeholder="quantity" name="quantity" className="form-control" value={this.state.quantity} onChange={this.changequantityHandler}/>

                                            <label> Reason:&emsp; </label>
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

                                            <label> Condition:&emsp; </label>
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

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateReturnItem}>Save</button>
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

export default CreateReturnItemComponent
