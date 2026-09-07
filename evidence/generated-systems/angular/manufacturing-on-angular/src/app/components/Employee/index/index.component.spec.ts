
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexEmployeeComponent } from './index.component';
import { EmployeeService } from '../../../services/Employee.service';

describe('IndexEmployeeComponent', () => {
  let component: IndexEmployeeComponent;
  let fixture: ComponentFixture<IndexEmployeeComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexEmployeeComponent
      ],
      providers: [
        EmployeeService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexEmployeeComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});