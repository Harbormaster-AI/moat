import React, { Component } from 'react'
import PositionService from '../services/PositionService';

class UpdatePositionComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                quantity: '',
                averageCost: '',
                marketValue: ''
        }
        this.updatePosition = this.updatePosition.bind(this);

        this.changequantityHandler = this.changequantityHandler.bind(this);
        this.changeaverageCostHandler = this.changeaverageCostHandler.bind(this);
        this.changemarketValueHandler = this.changemarketValueHandler.bind(this);
    }

    componentDidMount(){
        PositionService.getPositionById(this.state.id).then( (res) =>{
            let position = res.data;
            this.setState({
                quantity: position.quantity,
                averageCost: position.averageCost,
                marketValue: position.marketValue
            });
        });
    }

    updatePosition = (e) => {
        e.preventDefault();
        let position = {
            positionId: this.state.id,
            quantity: this.state.quantity,
            averageCost: this.state.averageCost,
            marketValue: this.state.marketValue
        };
        console.log('position => ' + JSON.stringify(position));
        console.log('id => ' + JSON.stringify(this.state.id));
        PositionService.updatePosition(position).then( res => {
            this.props.history.push('/positions');
        });
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

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update Position</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> quantity: </label>
                                                <input placeholder="quantity" name="quantity" className="form-control" value={this.state.quantity} onChange={this.changequantityHandler}/>

                                            <label> averageCost: </label>
                                                <input placeholder="averageCost" name="averageCost" className="form-control" value={this.state.averageCost} onChange={this.changeaverageCostHandler}/>

                                            <label> marketValue: </label>
                                                <input placeholder="marketValue" name="marketValue" className="form-control" value={this.state.marketValue} onChange={this.changemarketValueHandler}/>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updatePosition}>Save</button>
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

export default UpdatePositionComponent
