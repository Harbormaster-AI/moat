import React, { Component } from 'react'
import RegulationService from '../services/RegulationService';

class UpdateRegulationComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                name: '',
                citation: '',
                jurisdiction: '',
                publicationUrl: ''
        }
        this.updateRegulation = this.updateRegulation.bind(this);

        this.changenameHandler = this.changenameHandler.bind(this);
        this.changecitationHandler = this.changecitationHandler.bind(this);
        this.changejurisdictionHandler = this.changejurisdictionHandler.bind(this);
        this.changepublicationUrlHandler = this.changepublicationUrlHandler.bind(this);
    }

    componentDidMount(){
        RegulationService.getRegulationById(this.state.id).then( (res) =>{
            let regulation = res.data;
            this.setState({
                name: regulation.name,
                citation: regulation.citation,
                jurisdiction: regulation.jurisdiction,
                publicationUrl: regulation.publicationUrl
            });
        });
    }

    updateRegulation = (e) => {
        e.preventDefault();
        let regulation = {
            regulationId: this.state.id,
            name: this.state.name,
            citation: this.state.citation,
            jurisdiction: this.state.jurisdiction,
            publicationUrl: this.state.publicationUrl
        };
        console.log('regulation => ' + JSON.stringify(regulation));
        console.log('id => ' + JSON.stringify(this.state.id));
        RegulationService.updateRegulation(regulation).then( res => {
            this.props.history.push('/regulations');
        });
    }

    changenameHandler= (event) => {
        this.setState({name: event.target.value});
    }
    changecitationHandler= (event) => {
        this.setState({citation: event.target.value});
    }
    changejurisdictionHandler= (event) => {
        this.setState({jurisdiction: event.target.value});
    }
    changepublicationUrlHandler= (event) => {
        this.setState({publicationUrl: event.target.value});
    }

    cancel(){
        this.props.history.push('/regulations');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update Regulation</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> name: </label>
                                                <input placeholder="name" name="name" className="form-control" value={this.state.name} onChange={this.changenameHandler}/>

                                            <label> citation: </label>
                                                <input placeholder="citation" name="citation" className="form-control" value={this.state.citation} onChange={this.changecitationHandler}/>

                                            <label> jurisdiction: </label>
                                                <input placeholder="jurisdiction" name="jurisdiction" className="form-control" value={this.state.jurisdiction} onChange={this.changejurisdictionHandler}/>

                                            <label> publicationUrl: </label>
                                                <input placeholder="publicationUrl" name="publicationUrl" className="form-control" value={this.state.publicationUrl} onChange={this.changepublicationUrlHandler}/>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateRegulation}>Save</button>
                                        <button className="btn btn-danger" onClick={this.cancel.bind(this)} style={{marginLeft: "10px"}}>Cancel</button>
                                    </form>
                                </div>
                            </div>
                        </div>

                   </div>
            </div>
        )
    }
}

export default UpdateRegulationComponent
