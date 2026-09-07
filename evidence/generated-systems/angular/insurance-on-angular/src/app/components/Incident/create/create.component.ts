import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { IncidentService } from '../../../services/Incident.service';
import { Incident } from '../../../models/Incident';
import { SubBaseComponent } from '../../Incident/sub.base.component';

@Component({
    selector: 'app-create-incident',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateIncidentComponent extends SubBaseComponent implements OnInit {

    title = 'Add Incident';

    incidentForm: FormGroup;
    incident: Incident;

    constructor( http: HttpClient,
        private incidentService: IncidentService,
        private fb: FormBuilder,
        private router: Router
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

    
    addIncident(location, description, Claim, InsuredObjects, IncidentType): void {
        this.incidentService
        .addIncident(location, description, Claim, InsuredObjects, IncidentType)
            .subscribe(() => {
                this.router.navigate(['/indexIncident']);
            });
    }

    ngOnInit(): void {
    }
}