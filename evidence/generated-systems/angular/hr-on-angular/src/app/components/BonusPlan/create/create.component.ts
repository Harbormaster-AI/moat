import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { BonusPlanService } from '../../../services/BonusPlan.service';
import { BonusPlan } from '../../../models/BonusPlan';
import { SubBaseComponent } from '../../BonusPlan/sub.base.component';

@Component({
    selector: 'app-create-bonusPlan',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateBonusPlanComponent extends SubBaseComponent implements OnInit {

    title = 'Add BonusPlan';

    bonusPlanForm: FormGroup;
    bonusPlan: BonusPlan;

    constructor( http: HttpClient,
        private bonusPlanService: BonusPlanService,
        private fb: FormBuilder,
        private router: Router
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

    
    addBonusPlan(name, targetPercentage, CompensationPackages): void {
        this.bonusPlanService
        .addBonusPlan(name, targetPercentage, CompensationPackages)
            .subscribe(() => {
                this.router.navigate(['/indexBonusPlan']);
            });
    }

    ngOnInit(): void {
    }
}