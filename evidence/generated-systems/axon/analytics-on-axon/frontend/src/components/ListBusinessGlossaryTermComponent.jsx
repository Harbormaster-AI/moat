import React, { Component } from 'react'
import BusinessGlossaryTermService from '../services/BusinessGlossaryTermService'

class ListBusinessGlossaryTermComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                businessGlossaryTerms: []
        }
        this.addBusinessGlossaryTerm = this.addBusinessGlossaryTerm.bind(this);
        this.editBusinessGlossaryTerm = this.editBusinessGlossaryTerm.bind(this);
        this.deleteBusinessGlossaryTerm = this.deleteBusinessGlossaryTerm.bind(this);
    }

    deleteBusinessGlossaryTerm(id){
        BusinessGlossaryTermService.deleteBusinessGlossaryTerm(id).then( res => {
            this.setState({businessGlossaryTerms: this.state.businessGlossaryTerms.filter(businessGlossaryTerm => businessGlossaryTerm.businessGlossaryTermId !== id)});
        });
    }
    viewBusinessGlossaryTerm(id){
        this.props.history.push(`/view-businessGlossaryTerm/${id}`);
    }
    editBusinessGlossaryTerm(id){
        this.props.history.push(`/add-businessGlossaryTerm/${id}`);
    }

    componentDidMount(){
        BusinessGlossaryTermService.getBusinessGlossaryTerms().then((res) => {
            this.setState({ businessGlossaryTerms: res.data});
        });
    }

    addBusinessGlossaryTerm(){
        this.props.history.push('/add-businessGlossaryTerm/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">BusinessGlossaryTerm List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addBusinessGlossaryTerm}> Add BusinessGlossaryTerm</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> Term </th>
                                    <th> Definition </th>
                                    <th> Steward </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.businessGlossaryTerms.map(
                                        businessGlossaryTerm => 
                                        <tr key = {businessGlossaryTerm.businessGlossaryTermId}>
                                             <td> { businessGlossaryTerm.term } </td>
                                             <td> { businessGlossaryTerm.definition } </td>
                                             <td> { businessGlossaryTerm.steward } </td>
                                             <td>
                                                 <button onClick={ () => this.editBusinessGlossaryTerm(businessGlossaryTerm.businessGlossaryTermId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteBusinessGlossaryTerm(businessGlossaryTerm.businessGlossaryTermId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewBusinessGlossaryTerm(businessGlossaryTerm.businessGlossaryTermId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListBusinessGlossaryTermComponent
