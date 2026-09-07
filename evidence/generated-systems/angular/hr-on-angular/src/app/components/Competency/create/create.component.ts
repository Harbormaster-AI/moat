import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { CompetencyService } from '../../../services/Competency.service';
import { Competency } from '../../../models/Competency';
import { SubBaseComponent } from '../../Competency/sub.base.component';

@Component({
    selector: 'app-create-competency',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateCompetencyComponent extends SubBaseComponent implements OnInit {

    title = 'Add Competency';

    competencyForm: FormGroup;
    competency: Competency;

    constructor( http: HttpClient,
        private competencyService: CompetencyService,
        private fb: FormBuilder,
        private router: Router
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

    
    addCompetency(name, category, JobProfiles, CompetencyRatings): void {
        this.competencyService
        .addCompetency(name, category, JobProfiles, CompetencyRatings)
            .subscribe(() => {
                this.router.navigate(['/indexCompetency']);
            });
    }

    ngOnInit(): void {
    }
}