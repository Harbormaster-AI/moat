import React, { Component } from 'react'
import BonusPlanService from '../services/BonusPlanService';

class CreateBonusPlanComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                name: '',
                targetPercentage: ''
        }
        this.changenameHandler = this.changenameHandler.bind(this);
        this.changetargetPercentageHandler = this.changetargetPercentageHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            BonusPlanService.getBonusPlanById(this.state.id).then( (res) =>{
                let bonusPlan = res.data;
                this.setState({
                    name: bonusPlan.name,
                    targetPercentage: bonusPlan.targetPercentage
                });
            });
        }        
    }
    saveOrUpdateBonusPlan = (e) => {
        e.preventDefault();
        let bonusPlan = {
                bonusPlanId: this.state.id,
                name: this.state.name,
                targetPercentage: this.state.targetPercentage
            };
        console.log('bonusPlan => ' + JSON.stringify(bonusPlan));

        // step 5
        if(this.state.id === '_add'){
            bonusPlan.bonusPlanId=''
            BonusPlanService.createBonusPlan(bonusPlan).then(res =>{
                this.props.history.push('/bonusPlans');
            });
        }else{
            BonusPlanService.updateBonusPlan(bonusPlan).then( res => {
                this.props.history.push('/bonusPlans');
            });
        }
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

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add BonusPlan</h3>
        }else{
            return <h3 className="text-center">Update BonusPlan</h3>
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
                                            <label> name:&emsp; </label>
                                                <input placeholder="name" name="name" className="form-control" value={this.state.name} onChange={this.changenameHandler}/>

                                            <label> targetPercentage:&emsp; </label>
                                                <input placeholder="targetPercentage" name="targetPercentage" className="form-control" value={this.state.targetPercentage} onChange={this.changetargetPercentageHandler}/>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateBonusPlan}>Save</button>
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

export default CreateBonusPlanComponent
