import React, { Component } from 'react'
import ThirdPartyAssessmentService from '../services/ThirdPartyAssessmentService'

class ListThirdPartyAssessmentComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                thirdPartyAssessments: []
        }
        this.addThirdPartyAssessment = this.addThirdPartyAssessment.bind(this);
        this.editThirdPartyAssessment = this.editThirdPartyAssessment.bind(this);
        this.deleteThirdPartyAssessment = this.deleteThirdPartyAssessment.bind(this);
    }

    deleteThirdPartyAssessment(id){
        ThirdPartyAssessmentService.deleteThirdPartyAssessment(id).then( res => {
            this.setState({thirdPartyAssessments: this.state.thirdPartyAssessments.filter(thirdPartyAssessment => thirdPartyAssessment.thirdPartyAssessmentId !== id)});
        });
    }
    viewThirdPartyAssessment(id){
        this.props.history.push(`/view-thirdPartyAssessment/${id}`);
    }
    editThirdPartyAssessment(id){
        this.props.history.push(`/add-thirdPartyAssessment/${id}`);
    }

    componentDidMount(){
        ThirdPartyAssessmentService.getThirdPartyAssessments().then((res) => {
            this.setState({ thirdPartyAssessments: res.data});
        });
    }

    addThirdPartyAssessment(){
        this.props.history.push('/add-thirdPartyAssessment/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">ThirdPartyAssessment List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addThirdPartyAssessment}> Add ThirdPartyAssessment</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> AssessmentDate </th>
                                    <th> Assessor </th>
                                    <th> AssessmentType </th>
                                    <th> Result </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.thirdPartyAssessments.map(
                                        thirdPartyAssessment => 
                                        <tr key = {thirdPartyAssessment.thirdPartyAssessmentId}>
                                             <td> { thirdPartyAssessment.assessmentDate } </td>
                                             <td> { thirdPartyAssessment.assessor } </td>
                                             <td> { thirdPartyAssessment.assessmentType } </td>
                                             <td> { thirdPartyAssessment.result } </td>
                                             <td>
                                                 <button onClick={ () => this.editThirdPartyAssessment(thirdPartyAssessment.thirdPartyAssessmentId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteThirdPartyAssessment(thirdPartyAssessment.thirdPartyAssessmentId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewThirdPartyAssessment(thirdPartyAssessment.thirdPartyAssessmentId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListThirdPartyAssessmentComponent
