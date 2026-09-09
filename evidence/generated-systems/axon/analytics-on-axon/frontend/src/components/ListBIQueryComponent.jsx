import React, { Component } from 'react'
import BIQueryService from '../services/BIQueryService'

class ListBIQueryComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                bIQuerys: []
        }
        this.addBIQuery = this.addBIQuery.bind(this);
        this.editBIQuery = this.editBIQuery.bind(this);
        this.deleteBIQuery = this.deleteBIQuery.bind(this);
    }

    deleteBIQuery(id){
        BIQueryService.deleteBIQuery(id).then( res => {
            this.setState({bIQuerys: this.state.bIQuerys.filter(bIQuery => bIQuery.bIQueryId !== id)});
        });
    }
    viewBIQuery(id){
        this.props.history.push(`/view-bIQuery/${id}`);
    }
    editBIQuery(id){
        this.props.history.push(`/add-bIQuery/${id}`);
    }

    componentDidMount(){
        BIQueryService.getBIQuerys().then((res) => {
            this.setState({ bIQuerys: res.data});
        });
    }

    addBIQuery(){
        this.props.history.push('/add-bIQuery/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">BIQuery List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addBIQuery}> Add BIQuery</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> Name </th>
                                    <th> Text </th>
                                    <th> Dialect </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.bIQuerys.map(
                                        bIQuery => 
                                        <tr key = {bIQuery.bIQueryId}>
                                             <td> { bIQuery.name } </td>
                                             <td> { bIQuery.text } </td>
                                             <td> { bIQuery.dialect } </td>
                                             <td>
                                                 <button onClick={ () => this.editBIQuery(bIQuery.bIQueryId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteBIQuery(bIQuery.bIQueryId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewBIQuery(bIQuery.bIQueryId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListBIQueryComponent
