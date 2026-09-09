import React, { Component } from 'react'
import PositionService from '../services/PositionService';

class CreatePositionComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                quantity: '',
                averageCost: '',
                marketValue: ''
        }
        this.changequantityHandler = this.changequantityHandler.bind(this);
        this.changeaverageCostHandler = this.changeaverageCostHandler.bind(this);
        this.changemarketValueHandler = this.changemarketValueHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            PositionService.getPositionById(this.state.id).then( (res) =>{
                let position = res.data;
                this.setState({
                    quantity: position.quantity,
                    averageCost: position.averageCost,
                    marketValue: position.marketValue
                });
            });
        }        
    }
    saveOrUpdatePosition = (e) => {
        e.preventDefault();
        let position = {
                positionId: this.state.id,
                quantity: this.state.quantity,
                averageCost: this.state.averageCost,
                marketValue: this.state.marketValue
            };
        console.log('position => ' + JSON.stringify(position));

        // step 5
        if(this.state.id === '_add'){
            position.positionId=''
            PositionService.createPosition(position).then(res =>{
                this.props.history.push('/positions');
            });
        }else{
            PositionService.updatePosition(position).then( res => {
                this.props.history.push('/positions');
            });
        }
    }
    
    changequantityHandler= (event) => {
        this.setState({quantity: event.target.value});
    }
    changeaverageCostHandler= (event) => {
        this.setState({averageCost: event.target.value});
    }
    changemarketValueHandler= (event) => {
        this.setState({marketValue: event.target.value});
    }

    cancel(){
        this.props.history.push('/positions');
    }

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add Position</h3>
        }else{
            return <h3 className="text-center">Update Position</h3>
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
                                                <input placeholder="quantity" name="quantity" className="form-control" value={this.state.quantity} onChange={this.changequantityHandler}/>

                                            <label> averageCost:&emsp; </label>
                                                <input placeholder="averageCost" name="averageCost" className="form-control" value={this.state.averageCost} onChange={this.changeaverageCostHandler}/>

                                            <label> marketValue:&emsp; </label>
                                                <input placeholder="marketValue" name="marketValue" className="form-control" value={this.state.marketValue} onChange={this.changemarketValueHandler}/>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdatePosition}>Save</button>
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

export default CreatePositionComponent
