import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { ProductionLineService } from '../../../services/ProductionLine.service';
import { SubBaseComponent } from '../../ProductionLine/sub.base.component';


@Component({
    selector: 'app-edit-productionLine',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditProductionLineComponent extends SubBaseComponent implements OnInit {

    title = 'Edit ProductionLine';

    productionLineForm: FormGroup;
    productionLine: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: ProductionLineService,
        private fb: FormBuilder
) {
        super(http);
        this.productionLineForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  name: ['', Validators.required],
      Plant: ['', ],
      WorkCenters: ['', ],
      LineType: ['', ]
        });
    }

    
    updateProductionLine(name, Plant, WorkCenters, LineType): void {
        this.route.params.subscribe((params) => {

                        this.service.updateProductionLine(name, Plant, WorkCenters, LineType, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexProductionLine']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getProductionLine(params['id']).subscribe(res => {
                this.productionLine = res;
            });
        });
    }
}