import React, { Component } from 'react'
import BusinessUnitService from '../services/BusinessUnitService'

class ListBusinessUnitComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                businessUnits: []
        }
        this.addBusinessUnit = this.addBusinessUnit.bind(this);
        this.editBusinessUnit = this.editBusinessUnit.bind(this);
        this.deleteBusinessUnit = this.deleteBusinessUnit.bind(this);
    }

    deleteBusinessUnit(id){
        BusinessUnitService.deleteBusinessUnit(id).then( res => {
            this.setState({businessUnits: this.state.businessUnits.filter(businessUnit => businessUnit.businessUnitId !== id)});
        });
    }
    viewBusinessUnit(id){
        this.props.history.push(`/view-businessUnit/${id}`);
    }
    editBusinessUnit(id){
        this.props.history.push(`/add-businessUnit/${id}`);
    }

    componentDidMount(){
        BusinessUnitService.getBusinessUnits().then((res) => {
            this.setState({ businessUnits: res.data});
        });
    }

    addBusinessUnit(){
        this.props.history.push('/add-businessUnit/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">BusinessUnit List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addBusinessUnit}> Add BusinessUnit</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> Name </th>
                                    <th> Leader </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.businessUnits.map(
                                        businessUnit => 
                                        <tr key = {businessUnit.businessUnitId}>
                                             <td> { businessUnit.name } </td>
                                             <td> { businessUnit.leader } </td>
                                             <td>
                                                 <button onClick={ () => this.editBusinessUnit(businessUnit.businessUnitId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteBusinessUnit(businessUnit.businessUnitId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewBusinessUnit(businessUnit.businessUnitId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListBusinessUnitComponent
