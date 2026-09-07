import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { CompetencyService } from '../../../services/Competency.service';
import { SubBaseComponent } from '../../Competency/sub.base.component';


@Component({
    selector: 'app-edit-competency',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditCompetencyComponent extends SubBaseComponent implements OnInit {

    title = 'Edit Competency';

    competencyForm: FormGroup;
    competency: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: CompetencyService,
        private fb: FormBuilder
) {
        super(http);
        this.competencyForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  name: ['', Validators.required],
      category: ['', Validators.required],
      JobProfiles: ['', ],
      CompetencyRatings: ['', ]
        });
    }

    
    updateCompetency(name, category, JobProfiles, CompetencyRatings): void {
        this.route.params.subscribe((params) => {

                        this.service.updateCompetency(name, category, JobProfiles, CompetencyRatings, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexCompetency']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getCompetency(params['id']).subscribe(res => {
                this.competency = res;
            });
        });
    }
}