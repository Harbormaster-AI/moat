import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { CareTeamService } from '../../../services/CareTeam.service';
import { SubBaseComponent } from '../../CareTeam/sub.base.component';


@Component({
    selector: 'app-edit-careTeam',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditCareTeamComponent extends SubBaseComponent implements OnInit {

    title = 'Edit CareTeam';

    careTeamForm: FormGroup;
    careTeam: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: CareTeamService,
        private fb: FormBuilder
) {
        super(http);
        this.careTeamForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  name: ['', Validators.required],
      Department: ['', ],
      Clinicians: ['', ],
      Patients: ['', ],
      CareSetting: ['', ]
        });
    }

    
    updateCareTeam(name, Department, Clinicians, Patients, CareSetting): void {
        this.route.params.subscribe((params) => {

                        this.service.updateCareTeam(name, Department, Clinicians, Patients, CareSetting, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexCareTeam']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getCareTeam(params['id']).subscribe(res => {
                this.careTeam = res;
            });
        });
    }
}