
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateRunParameterComponent } from './create.component';
import { RunParameterService } from '../../../services/RunParameter.service';
import { Router } from '@angular/router';

describe('CreateRunParameterComponent', () => {
  let component: CreateRunParameterComponent;
  let fixture: ComponentFixture<CreateRunParameterComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateRunParameterComponent
      ],
      providers: [
        RunParameterService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateRunParameterComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});