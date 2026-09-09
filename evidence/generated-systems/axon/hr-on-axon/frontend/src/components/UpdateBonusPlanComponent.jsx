import React, { Component } from 'react'
import BonusPlanService from '../services/BonusPlanService';

class UpdateBonusPlanComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                name: '',
                targetPercentage: ''
        }
        this.updateBonusPlan = this.updateBonusPlan.bind(this);

        this.changenameHandler = this.changenameHandler.bind(this);
        this.changetargetPercentageHandler = this.changetargetPercentageHandler.bind(this);
    }

    componentDidMount(){
        BonusPlanService.getBonusPlanById(this.state.id).then( (res) =>{
            let bonusPlan = res.data;
            this.setState({
                name: bonusPlan.name,
                targetPercentage: bonusPlan.targetPercentage
            });
        });
    }

    updateBonusPlan = (e) => {
        e.preventDefault();
        let bonusPlan = {
            bonusPlanId: this.state.id,
            name: this.state.name,
            targetPercentage: this.state.targetPercentage
        };
        console.log('bonusPlan => ' + JSON.stringify(bonusPlan));
        console.log('id => ' + JSON.stringify(this.state.id));
        BonusPlanService.updateBonusPlan(bonusPlan).then( res => {
            this.props.history.push('/bonusPlans');
        });
    }

    changenameHandler= (event) => {
        this.setState({name: event.target.value});
    }
    changetargetPercentageHandler= (event) => {
        this.setState({targetPercentage: event.target.value});
    }

    cancel(){
        this.props.history.push('/bonusPlans');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update BonusPlan</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> name: </label>
                                                <input placeholder="name" name="name" className="form-control" value={this.state.name} onChange={this.changenameHandler}/>

                                            <label> targetPercentage: </label>
                                                <input placeholder="targetPercentage" name="targetPercentage" className="form-control" value={this.state.targetPercentage} onChange={this.changetargetPercentageHandler}/>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateBonusPlan}>Save</button>
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

export default UpdateBonusPlanComponent
