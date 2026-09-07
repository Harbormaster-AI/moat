
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateMRPRunComponent } from './create.component';
import { MRPRunService } from '../../../services/MRPRun.service';
import { Router } from '@angular/router';

describe('CreateMRPRunComponent', () => {
  let component: CreateMRPRunComponent;
  let fixture: ComponentFixture<CreateMRPRunComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateMRPRunComponent
      ],
      providers: [
        MRPRunService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateMRPRunComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});