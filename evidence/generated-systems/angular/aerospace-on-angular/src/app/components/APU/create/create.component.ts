import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { APUService } from '../../../services/APU.service';
import { APU } from '../../../models/APU';
import { SubBaseComponent } from '../../APU/sub.base.component';

@Component({
    selector: 'app-create-aPU',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateAPUComponent extends SubBaseComponent implements OnInit {

    title = 'Add APU';

    aPUForm: FormGroup;
    aPU: APU;

    constructor( http: HttpClient,
        private aPUService: APUService,
        private fb: FormBuilder,
        private router: Router
) {
        super(http);
        this.aPUForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  model_: ['', Validators.required],
      Supplier: ['', ],
      Variants: ['', ]
        });
    }

    
    addAPU(model_, Supplier, Variants): void {
        this.aPUService
        .addAPU(model_, Supplier, Variants)
            .subscribe(() => {
                this.router.navigate(['/indexAPU']);
            });
    }

    ngOnInit(): void {
    }
}