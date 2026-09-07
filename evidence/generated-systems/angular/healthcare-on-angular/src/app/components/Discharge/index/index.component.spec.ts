
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexDischargeComponent } from './index.component';
import { DischargeService } from '../../../services/Discharge.service';

describe('IndexDischargeComponent', () => {
  let component: IndexDischargeComponent;
  let fixture: ComponentFixture<IndexDischargeComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexDischargeComponent
      ],
      providers: [
        DischargeService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexDischargeComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});