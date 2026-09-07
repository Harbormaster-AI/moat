import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { BonusPlanService } from '../../../services/BonusPlan.service';
import { SubBaseComponent } from '../../BonusPlan/sub.base.component';


@Component({
    selector: 'app-edit-bonusPlan',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditBonusPlanComponent extends SubBaseComponent implements OnInit {

    title = 'Edit BonusPlan';

    bonusPlanForm: FormGroup;
    bonusPlan: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: BonusPlanService,
        private fb: FormBuilder
) {
        super(http);
        this.bonusPlanForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  name: ['', Validators.required],
      targetPercentage: ['', Validators.required],
      CompensationPackages: ['', ]
        });
    }

    
    updateBonusPlan(name, targetPercentage, CompensationPackages): void {
        this.route.params.subscribe((params) => {

                        this.service.updateBonusPlan(name, targetPercentage, CompensationPackages, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexBonusPlan']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getBonusPlan(params['id']).subscribe(res => {
                this.bonusPlan = res;
            });
        });
    }
}