import React, { Component } from 'react'
import AirworthinessDirectiveService from '../services/AirworthinessDirectiveService'

class ListAirworthinessDirectiveComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                airworthinessDirectives: []
        }
        this.addAirworthinessDirective = this.addAirworthinessDirective.bind(this);
        this.editAirworthinessDirective = this.editAirworthinessDirective.bind(this);
        this.deleteAirworthinessDirective = this.deleteAirworthinessDirective.bind(this);
    }

    deleteAirworthinessDirective(id){
        AirworthinessDirectiveService.deleteAirworthinessDirective(id).then( res => {
            this.setState({airworthinessDirectives: this.state.airworthinessDirectives.filter(airworthinessDirective => airworthinessDirective.airworthinessDirectiveId !== id)});
        });
    }
    viewAirworthinessDirective(id){
        this.props.history.push(`/view-airworthinessDirective/${id}`);
    }
    editAirworthinessDirective(id){
        this.props.history.push(`/add-airworthinessDirective/${id}`);
    }

    componentDidMount(){
        AirworthinessDirectiveService.getAirworthinessDirectives().then((res) => {
            this.setState({ airworthinessDirectives: res.data});
        });
    }

    addAirworthinessDirective(){
        this.props.history.push('/add-airworthinessDirective/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">AirworthinessDirective List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addAirworthinessDirective}> Add AirworthinessDirective</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> DirectiveNumber </th>
                                    <th> Title </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.airworthinessDirectives.map(
                                        airworthinessDirective => 
                                        <tr key = {airworthinessDirective.airworthinessDirectiveId}>
                                             <td> { airworthinessDirective.directiveNumber } </td>
                                             <td> { airworthinessDirective.title } </td>
                                             <td>
                                                 <button onClick={ () => this.editAirworthinessDirective(airworthinessDirective.airworthinessDirectiveId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteAirworthinessDirective(airworthinessDirective.airworthinessDirectiveId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewAirworthinessDirective(airworthinessDirective.airworthinessDirectiveId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListAirworthinessDirectiveComponent
