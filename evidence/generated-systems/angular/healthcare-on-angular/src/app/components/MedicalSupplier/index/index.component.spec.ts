
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexMedicalSupplierComponent } from './index.component';
import { MedicalSupplierService } from '../../../services/MedicalSupplier.service';

describe('IndexMedicalSupplierComponent', () => {
  let component: IndexMedicalSupplierComponent;
  let fixture: ComponentFixture<IndexMedicalSupplierComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexMedicalSupplierComponent
      ],
      providers: [
        MedicalSupplierService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexMedicalSupplierComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});