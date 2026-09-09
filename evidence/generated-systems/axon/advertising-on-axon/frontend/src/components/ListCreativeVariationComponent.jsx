import React, { Component } from 'react'
import CreativeVariationService from '../services/CreativeVariationService'

class ListCreativeVariationComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                creativeVariations: []
        }
        this.addCreativeVariation = this.addCreativeVariation.bind(this);
        this.editCreativeVariation = this.editCreativeVariation.bind(this);
        this.deleteCreativeVariation = this.deleteCreativeVariation.bind(this);
    }

    deleteCreativeVariation(id){
        CreativeVariationService.deleteCreativeVariation(id).then( res => {
            this.setState({creativeVariations: this.state.creativeVariations.filter(creativeVariation => creativeVariation.creativeVariationId !== id)});
        });
    }
    viewCreativeVariation(id){
        this.props.history.push(`/view-creativeVariation/${id}`);
    }
    editCreativeVariation(id){
        this.props.history.push(`/add-creativeVariation/${id}`);
    }

    componentDidMount(){
        CreativeVariationService.getCreativeVariations().then((res) => {
            this.setState({ creativeVariations: res.data});
        });
    }

    addCreativeVariation(){
        this.props.history.push('/add-creativeVariation/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">CreativeVariation List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addCreativeVariation}> Add CreativeVariation</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> Name </th>
                                    <th> Language </th>
                                    <th> Headline </th>
                                    <th> BodyText </th>
                                    <th> CallToAction </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.creativeVariations.map(
                                        creativeVariation => 
                                        <tr key = {creativeVariation.creativeVariationId}>
                                             <td> { creativeVariation.name } </td>
                                             <td> { creativeVariation.language } </td>
                                             <td> { creativeVariation.headline } </td>
                                             <td> { creativeVariation.bodyText } </td>
                                             <td> { creativeVariation.callToAction } </td>
                                             <td>
                                                 <button onClick={ () => this.editCreativeVariation(creativeVariation.creativeVariationId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteCreativeVariation(creativeVariation.creativeVariationId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewCreativeVariation(creativeVariation.creativeVariationId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListCreativeVariationComponent
