import React, { Component } from 'react'
import RegulationService from '../services/RegulationService'

class ListRegulationComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                regulations: []
        }
        this.addRegulation = this.addRegulation.bind(this);
        this.editRegulation = this.editRegulation.bind(this);
        this.deleteRegulation = this.deleteRegulation.bind(this);
    }

    deleteRegulation(id){
        RegulationService.deleteRegulation(id).then( res => {
            this.setState({regulations: this.state.regulations.filter(regulation => regulation.regulationId !== id)});
        });
    }
    viewRegulation(id){
        this.props.history.push(`/view-regulation/${id}`);
    }
    editRegulation(id){
        this.props.history.push(`/add-regulation/${id}`);
    }

    componentDidMount(){
        RegulationService.getRegulations().then((res) => {
            this.setState({ regulations: res.data});
        });
    }

    addRegulation(){
        this.props.history.push('/add-regulation/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">Regulation List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addRegulation}> Add Regulation</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> Name </th>
                                    <th> Citation </th>
                                    <th> Jurisdiction </th>
                                    <th> PublicationUrl </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.regulations.map(
                                        regulation => 
                                        <tr key = {regulation.regulationId}>
                                             <td> { regulation.name } </td>
                                             <td> { regulation.citation } </td>
                                             <td> { regulation.jurisdiction } </td>
                                             <td> { regulation.publicationUrl } </td>
                                             <td>
                                                 <button onClick={ () => this.editRegulation(regulation.regulationId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteRegulation(regulation.regulationId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewRegulation(regulation.regulationId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListRegulationComponent
