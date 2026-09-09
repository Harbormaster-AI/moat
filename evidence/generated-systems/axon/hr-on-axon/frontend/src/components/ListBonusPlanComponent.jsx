import React, { Component } from 'react'
import BonusPlanService from '../services/BonusPlanService'

class ListBonusPlanComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                bonusPlans: []
        }
        this.addBonusPlan = this.addBonusPlan.bind(this);
        this.editBonusPlan = this.editBonusPlan.bind(this);
        this.deleteBonusPlan = this.deleteBonusPlan.bind(this);
    }

    deleteBonusPlan(id){
        BonusPlanService.deleteBonusPlan(id).then( res => {
            this.setState({bonusPlans: this.state.bonusPlans.filter(bonusPlan => bonusPlan.bonusPlanId !== id)});
        });
    }
    viewBonusPlan(id){
        this.props.history.push(`/view-bonusPlan/${id}`);
    }
    editBonusPlan(id){
        this.props.history.push(`/add-bonusPlan/${id}`);
    }

    componentDidMount(){
        BonusPlanService.getBonusPlans().then((res) => {
            this.setState({ bonusPlans: res.data});
        });
    }

    addBonusPlan(){
        this.props.history.push('/add-bonusPlan/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">BonusPlan List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addBonusPlan}> Add BonusPlan</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> Name </th>
                                    <th> TargetPercentage </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.bonusPlans.map(
                                        bonusPlan => 
                                        <tr key = {bonusPlan.bonusPlanId}>
                                             <td> { bonusPlan.name } </td>
                                             <td> { bonusPlan.targetPercentage } </td>
                                             <td>
                                                 <button onClick={ () => this.editBonusPlan(bonusPlan.bonusPlanId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteBonusPlan(bonusPlan.bonusPlanId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewBonusPlan(bonusPlan.bonusPlanId)} className="btn btn-outline-info btn-sm">View </button>
                                             </td>
                                        </tr>
                                    )
                                }
                            </tbody>
                        </table>

                 </div>

            </div>
        )
    }
}

export default ListBonusPlanComponent
