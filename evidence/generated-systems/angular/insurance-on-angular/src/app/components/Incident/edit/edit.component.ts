import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { IncidentService } from '../../../services/Incident.service';
import { SubBaseComponent } from '../../Incident/sub.base.component';


@Component({
    selector: 'app-edit-incident',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditIncidentComponent extends SubBaseComponent implements OnInit {

    title = 'Edit Incident';

    incidentForm: FormGroup;
    incident: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: IncidentService,
        private fb: FormBuilder
) {
        super(http);
        this.incidentForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  location: ['', Validators.required],
      description: ['', Validators.required],
      Claim: ['', ],
      InsuredObjects: ['', ],
      IncidentType: ['', ]
        });
    }

    
    updateIncident(location, description, Claim, InsuredObjects, IncidentType): void {
        this.route.params.subscribe((params) => {

                        this.service.updateIncident(location, description, Claim, InsuredObjects, IncidentType, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexIncident']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getIncident(params['id']).subscribe(res => {
                this.incident = res;
            });
        });
    }
}