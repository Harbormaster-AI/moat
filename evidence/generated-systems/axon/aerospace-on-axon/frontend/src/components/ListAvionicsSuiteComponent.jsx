import React, { Component } from 'react'
import AvionicsSuiteService from '../services/AvionicsSuiteService'

class ListAvionicsSuiteComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                avionicsSuites: []
        }
        this.addAvionicsSuite = this.addAvionicsSuite.bind(this);
        this.editAvionicsSuite = this.editAvionicsSuite.bind(this);
        this.deleteAvionicsSuite = this.deleteAvionicsSuite.bind(this);
    }

    deleteAvionicsSuite(id){
        AvionicsSuiteService.deleteAvionicsSuite(id).then( res => {
            this.setState({avionicsSuites: this.state.avionicsSuites.filter(avionicsSuite => avionicsSuite.avionicsSuiteId !== id)});
        });
    }
    viewAvionicsSuite(id){
        this.props.history.push(`/view-avionicsSuite/${id}`);
    }
    editAvionicsSuite(id){
        this.props.history.push(`/add-avionicsSuite/${id}`);
    }

    componentDidMount(){
        AvionicsSuiteService.getAvionicsSuites().then((res) => {
            this.setState({ avionicsSuites: res.data});
        });
    }

    addAvionicsSuite(){
        this.props.history.push('/add-avionicsSuite/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">AvionicsSuite List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addAvionicsSuite}> Add AvionicsSuite</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> SuiteName </th>
                                    <th> SoftwareBaseline </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.avionicsSuites.map(
                                        avionicsSuite => 
                                        <tr key = {avionicsSuite.avionicsSuiteId}>
                                             <td> { avionicsSuite.suiteName } </td>
                                             <td> { avionicsSuite.softwareBaseline } </td>
                                             <td>
                                                 <button onClick={ () => this.editAvionicsSuite(avionicsSuite.avionicsSuiteId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteAvionicsSuite(avionicsSuite.avionicsSuiteId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewAvionicsSuite(avionicsSuite.avionicsSuiteId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListAvionicsSuiteComponent
