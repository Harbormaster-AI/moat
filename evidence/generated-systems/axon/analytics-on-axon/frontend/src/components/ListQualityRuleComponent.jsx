import React, { Component } from 'react'
import QualityRuleService from '../services/QualityRuleService'

class ListQualityRuleComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                qualityRules: []
        }
        this.addQualityRule = this.addQualityRule.bind(this);
        this.editQualityRule = this.editQualityRule.bind(this);
        this.deleteQualityRule = this.deleteQualityRule.bind(this);
    }

    deleteQualityRule(id){
        QualityRuleService.deleteQualityRule(id).then( res => {
            this.setState({qualityRules: this.state.qualityRules.filter(qualityRule => qualityRule.qualityRuleId !== id)});
        });
    }
    viewQualityRule(id){
        this.props.history.push(`/view-qualityRule/${id}`);
    }
    editQualityRule(id){
        this.props.history.push(`/add-qualityRule/${id}`);
    }

    componentDidMount(){
        QualityRuleService.getQualityRules().then((res) => {
            this.setState({ qualityRules: res.data});
        });
    }

    addQualityRule(){
        this.props.history.push('/add-qualityRule/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">QualityRule List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addQualityRule}> Add QualityRule</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> Name </th>
                                    <th> Threshold </th>
                                    <th> TargetField </th>
                                    <th> Dimension </th>
                                    <th> Operator </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.qualityRules.map(
                                        qualityRule => 
                                        <tr key = {qualityRule.qualityRuleId}>
                                             <td> { qualityRule.name } </td>
                                             <td> { qualityRule.threshold } </td>
                                             <td> { qualityRule.targetField } </td>
                                             <td> { qualityRule.dimension } </td>
                                             <td> { qualityRule.operator } </td>
                                             <td>
                                                 <button onClick={ () => this.editQualityRule(qualityRule.qualityRuleId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteQualityRule(qualityRule.qualityRuleId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewQualityRule(qualityRule.qualityRuleId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListQualityRuleComponent
