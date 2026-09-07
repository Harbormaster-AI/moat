import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { ConnectedAircraftService } from '../../../services/ConnectedAircraft.service';
import { SubBaseComponent } from '../../ConnectedAircraft/sub.base.component';


@Component({
    selector: 'app-edit-connectedAircraft',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditConnectedAircraftComponent extends SubBaseComponent implements OnInit {

    title = 'Edit ConnectedAircraft';

    connectedAircraftForm: FormGroup;
    connectedAircraft: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: ConnectedAircraftService,
        private fb: FormBuilder
) {
        super(http);
        this.connectedAircraftForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  communicationsProvider: ['', Validators.required],
      Aircraft: ['', ],
      FlightHealthEvents: ['', ],
      SoftwareLoads: ['', ],
      ConnectivityStatus: ['', ]
        });
    }

    
    updateConnectedAircraft(communicationsProvider, Aircraft, FlightHealthEvents, SoftwareLoads, ConnectivityStatus): void {
        this.route.params.subscribe((params) => {

                        this.service.updateConnectedAircraft(communicationsProvider, Aircraft, FlightHealthEvents, SoftwareLoads, ConnectivityStatus, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexConnectedAircraft']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getConnectedAircraft(params['id']).subscribe(res => {
                this.connectedAircraft = res;
            });
        });
    }
}