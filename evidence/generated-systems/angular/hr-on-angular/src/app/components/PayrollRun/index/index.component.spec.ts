
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexPayrollRunComponent } from './index.component';
import { PayrollRunService } from '../../../services/PayrollRun.service';

describe('IndexPayrollRunComponent', () => {
  let component: IndexPayrollRunComponent;
  let fixture: ComponentFixture<IndexPayrollRunComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexPayrollRunComponent
      ],
      providers: [
        PayrollRunService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexPayrollRunComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});