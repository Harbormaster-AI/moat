
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateControlComponent } from './create.component';
import { ControlService } from '../../../services/Control.service';
import { Router } from '@angular/router';

describe('CreateControlComponent', () => {
  let component: CreateControlComponent;
  let fixture: ComponentFixture<CreateControlComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateControlComponent
      ],
      providers: [
        ControlService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateControlComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});