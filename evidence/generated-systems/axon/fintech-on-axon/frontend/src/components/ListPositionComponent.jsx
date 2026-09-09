import React, { Component } from 'react'
import PositionService from '../services/PositionService'

class ListPositionComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                positions: []
        }
        this.addPosition = this.addPosition.bind(this);
        this.editPosition = this.editPosition.bind(this);
        this.deletePosition = this.deletePosition.bind(this);
    }

    deletePosition(id){
        PositionService.deletePosition(id).then( res => {
            this.setState({positions: this.state.positions.filter(position => position.positionId !== id)});
        });
    }
    viewPosition(id){
        this.props.history.push(`/view-position/${id}`);
    }
    editPosition(id){
        this.props.history.push(`/add-position/${id}`);
    }

    componentDidMount(){
        PositionService.getPositions().then((res) => {
            this.setState({ positions: res.data});
        });
    }

    addPosition(){
        this.props.history.push('/add-position/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">Position List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addPosition}> Add Position</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> Quantity </th>
                                    <th> AverageCost </th>
                                    <th> MarketValue </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.positions.map(
                                        position => 
                                        <tr key = {position.positionId}>
                                             <td> { position.quantity } </td>
                                             <td> { position.averageCost } </td>
                                             <td> { position.marketValue } </td>
                                             <td>
                                                 <button onClick={ () => this.editPosition(position.positionId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deletePosition(position.positionId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewPosition(position.positionId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListPositionComponent
