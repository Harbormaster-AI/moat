import React, { Component } from 'react'
import QualitySpecificationService from '../services/QualitySpecificationService'

class ListQualitySpecificationComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                qualitySpecifications: []
        }
        this.addQualitySpecification = this.addQualitySpecification.bind(this);
        this.editQualitySpecification = this.editQualitySpecification.bind(this);
        this.deleteQualitySpecification = this.deleteQualitySpecification.bind(this);
    }

    deleteQualitySpecification(id){
        QualitySpecificationService.deleteQualitySpecification(id).then( res => {
            this.setState({qualitySpecifications: this.state.qualitySpecifications.filter(qualitySpecification => qualitySpecification.qualitySpecificationId !== id)});
        });
    }
    viewQualitySpecification(id){
        this.props.history.push(`/view-qualitySpecification/${id}`);
    }
    editQualitySpecification(id){
        this.props.history.push(`/add-qualitySpecification/${id}`);
    }

    componentDidMount(){
        QualitySpecificationService.getQualitySpecifications().then((res) => {
            this.setState({ qualitySpecifications: res.data});
        });
    }

    addQualitySpecification(){
        this.props.history.push('/add-qualitySpecification/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">QualitySpecification List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addQualitySpecification}> Add QualitySpecification</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> SpecCode </th>
                                    <th> Name </th>
                                    <th> Version </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.qualitySpecifications.map(
                                        qualitySpecification => 
                                        <tr key = {qualitySpecification.qualitySpecificationId}>
                                             <td> { qualitySpecification.specCode } </td>
                                             <td> { qualitySpecification.name } </td>
                                             <td> { qualitySpecification.version } </td>
                                             <td>
                                                 <button onClick={ () => this.editQualitySpecification(qualitySpecification.qualitySpecificationId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteQualitySpecification(qualitySpecification.qualitySpecificationId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewQualitySpecification(qualitySpecification.qualitySpecificationId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListQualitySpecificationComponent
