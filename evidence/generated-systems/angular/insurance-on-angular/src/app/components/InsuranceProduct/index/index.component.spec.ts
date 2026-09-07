
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexInsuranceProductComponent } from './index.component';
import { InsuranceProductService } from '../../../services/InsuranceProduct.service';

describe('IndexInsuranceProductComponent', () => {
  let component: IndexInsuranceProductComponent;
  let fixture: ComponentFixture<IndexInsuranceProductComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexInsuranceProductComponent
      ],
      providers: [
        InsuranceProductService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexInsuranceProductComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});