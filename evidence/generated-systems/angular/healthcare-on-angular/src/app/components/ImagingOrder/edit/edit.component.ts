import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { ImagingOrderService } from '../../../services/ImagingOrder.service';
import { SubBaseComponent } from '../../ImagingOrder/sub.base.component';


@Component({
    selector: 'app-edit-imagingOrder',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditImagingOrderComponent extends SubBaseComponent implements OnInit {

    title = 'Edit ImagingOrder';

    imagingOrderForm: FormGroup;
    imagingOrder: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: ImagingOrderService,
        private fb: FormBuilder
) {
        super(http);
        this.imagingOrderForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  bodySite: ['', Validators.required],
      contrast: ['', Validators.required],
      Order: ['', ],
      ImagingCenter: ['', ],
      Reports: ['', ],
      Modality: ['', ]
        });
    }

    
    updateImagingOrder(bodySite, contrast, Order, ImagingCenter, Reports, Modality): void {
        this.route.params.subscribe((params) => {

                        this.service.updateImagingOrder(bodySite, contrast, Order, ImagingCenter, Reports, Modality, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexImagingOrder']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getImagingOrder(params['id']).subscribe(res => {
                this.imagingOrder = res;
            });
        });
    }
}