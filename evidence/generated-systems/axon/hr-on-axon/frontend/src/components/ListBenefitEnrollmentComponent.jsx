import React, { Component } from 'react'
import BenefitEnrollmentService from '../services/BenefitEnrollmentService'

class ListBenefitEnrollmentComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                benefitEnrollments: []
        }
        this.addBenefitEnrollment = this.addBenefitEnrollment.bind(this);
        this.editBenefitEnrollment = this.editBenefitEnrollment.bind(this);
        this.deleteBenefitEnrollment = this.deleteBenefitEnrollment.bind(this);
    }

    deleteBenefitEnrollment(id){
        BenefitEnrollmentService.deleteBenefitEnrollment(id).then( res => {
            this.setState({benefitEnrollments: this.state.benefitEnrollments.filter(benefitEnrollment => benefitEnrollment.benefitEnrollmentId !== id)});
        });
    }
    viewBenefitEnrollment(id){
        this.props.history.push(`/view-benefitEnrollment/${id}`);
    }
    editBenefitEnrollment(id){
        this.props.history.push(`/add-benefitEnrollment/${id}`);
    }

    componentDidMount(){
        BenefitEnrollmentService.getBenefitEnrollments().then((res) => {
            this.setState({ benefitEnrollments: res.data});
        });
    }

    addBenefitEnrollment(){
        this.props.history.push('/add-benefitEnrollment/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">BenefitEnrollment List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addBenefitEnrollment}> Add BenefitEnrollment</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> EnrollmentId </th>
                                    <th> EffectiveFrom </th>
                                    <th> EffectiveTo </th>
                                    <th> Status </th>
                                    <th> CoverageLevel </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.benefitEnrollments.map(
                                        benefitEnrollment => 
                                        <tr key = {benefitEnrollment.benefitEnrollmentId}>
                                             <td> { benefitEnrollment.enrollmentId } </td>
                                             <td> { benefitEnrollment.effectiveFrom } </td>
                                             <td> { benefitEnrollment.effectiveTo } </td>
                                             <td> { benefitEnrollment.status } </td>
                                             <td> { benefitEnrollment.coverageLevel } </td>
                                             <td>
                                                 <button onClick={ () => this.editBenefitEnrollment(benefitEnrollment.benefitEnrollmentId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteBenefitEnrollment(benefitEnrollment.benefitEnrollmentId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewBenefitEnrollment(benefitEnrollment.benefitEnrollmentId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListBenefitEnrollmentComponent
