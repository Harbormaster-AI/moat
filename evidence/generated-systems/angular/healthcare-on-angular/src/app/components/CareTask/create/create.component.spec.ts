
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateCareTaskComponent } from './create.component';
import { CareTaskService } from '../../../services/CareTask.service';
import { Router } from '@angular/router';

describe('CreateCareTaskComponent', () => {
  let component: CreateCareTaskComponent;
  let fixture: ComponentFixture<CreateCareTaskComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateCareTaskComponent
      ],
      providers: [
        CareTaskService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateCareTaskComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});