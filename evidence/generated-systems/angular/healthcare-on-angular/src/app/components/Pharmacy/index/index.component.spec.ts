
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexPharmacyComponent } from './index.component';
import { PharmacyService } from '../../../services/Pharmacy.service';

describe('IndexPharmacyComponent', () => {
  let component: IndexPharmacyComponent;
  let fixture: ComponentFixture<IndexPharmacyComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexPharmacyComponent
      ],
      providers: [
        PharmacyService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexPharmacyComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});