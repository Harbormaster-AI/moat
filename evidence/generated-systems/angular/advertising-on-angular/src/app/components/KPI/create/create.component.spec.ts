
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateKPIComponent } from './create.component';
import { KPIService } from '../../../services/KPI.service';
import { Router } from '@angular/router';

describe('CreateKPIComponent', () => {
  let component: CreateKPIComponent;
  let fixture: ComponentFixture<CreateKPIComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateKPIComponent
      ],
      providers: [
        KPIService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateKPIComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});