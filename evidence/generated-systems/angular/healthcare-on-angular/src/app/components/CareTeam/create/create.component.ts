import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { CareTeamService } from '../../../services/CareTeam.service';
import { CareTeam } from '../../../models/CareTeam';
import { SubBaseComponent } from '../../CareTeam/sub.base.component';

@Component({
    selector: 'app-create-careTeam',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateCareTeamComponent extends SubBaseComponent implements OnInit {

    title = 'Add CareTeam';

    careTeamForm: FormGroup;
    careTeam: CareTeam;

    constructor( http: HttpClient,
        private careTeamService: CareTeamService,
        private fb: FormBuilder,
        private router: Router
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

    
    addCareTeam(name, Department, Clinicians, Patients, CareSetting): void {
        this.careTeamService
        .addCareTeam(name, Department, Clinicians, Patients, CareSetting)
            .subscribe(() => {
                this.router.navigate(['/indexCareTeam']);
            });
    }

    ngOnInit(): void {
    }
}