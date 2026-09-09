import React, { Component } from 'react'
import CabinLayoutService from '../services/CabinLayoutService'

class ListCabinLayoutComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                cabinLayouts: []
        }
        this.addCabinLayout = this.addCabinLayout.bind(this);
        this.editCabinLayout = this.editCabinLayout.bind(this);
        this.deleteCabinLayout = this.deleteCabinLayout.bind(this);
    }

    deleteCabinLayout(id){
        CabinLayoutService.deleteCabinLayout(id).then( res => {
            this.setState({cabinLayouts: this.state.cabinLayouts.filter(cabinLayout => cabinLayout.cabinLayoutId !== id)});
        });
    }
    viewCabinLayout(id){
        this.props.history.push(`/view-cabinLayout/${id}`);
    }
    editCabinLayout(id){
        this.props.history.push(`/add-cabinLayout/${id}`);
    }

    componentDidMount(){
        CabinLayoutService.getCabinLayouts().then((res) => {
            this.setState({ cabinLayouts: res.data});
        });
    }

    addCabinLayout(){
        this.props.history.push('/add-cabinLayout/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">CabinLayout List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addCabinLayout}> Add CabinLayout</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> LayoutCode </th>
                                    <th> TotalSeats </th>
                                    <th> ClassLayout </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.cabinLayouts.map(
                                        cabinLayout => 
                                        <tr key = {cabinLayout.cabinLayoutId}>
                                             <td> { cabinLayout.layoutCode } </td>
                                             <td> { cabinLayout.totalSeats } </td>
                                             <td> { cabinLayout.classLayout } </td>
                                             <td>
                                                 <button onClick={ () => this.editCabinLayout(cabinLayout.cabinLayoutId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteCabinLayout(cabinLayout.cabinLayoutId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewCabinLayout(cabinLayout.cabinLayoutId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListCabinLayoutComponent
