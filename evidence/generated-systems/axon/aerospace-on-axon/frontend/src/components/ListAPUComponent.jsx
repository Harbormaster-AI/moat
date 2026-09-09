import React, { Component } from 'react'
import APUService from '../services/APUService'

class ListAPUComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                aPUs: []
        }
        this.addAPU = this.addAPU.bind(this);
        this.editAPU = this.editAPU.bind(this);
        this.deleteAPU = this.deleteAPU.bind(this);
    }

    deleteAPU(id){
        APUService.deleteAPU(id).then( res => {
            this.setState({aPUs: this.state.aPUs.filter(aPU => aPU.aPUId !== id)});
        });
    }
    viewAPU(id){
        this.props.history.push(`/view-aPU/${id}`);
    }
    editAPU(id){
        this.props.history.push(`/add-aPU/${id}`);
    }

    componentDidMount(){
        APUService.getAPUs().then((res) => {
            this.setState({ aPUs: res.data});
        });
    }

    addAPU(){
        this.props.history.push('/add-aPU/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">APU List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addAPU}> Add APU</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> Model </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.aPUs.map(
                                        aPU => 
                                        <tr key = {aPU.aPUId}>
                                             <td> { aPU.model } </td>
                                             <td>
                                                 <button onClick={ () => this.editAPU(aPU.aPUId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteAPU(aPU.aPUId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewAPU(aPU.aPUId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListAPUComponent
