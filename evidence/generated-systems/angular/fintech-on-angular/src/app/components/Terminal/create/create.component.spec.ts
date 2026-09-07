
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateTerminalComponent } from './create.component';
import { TerminalService } from '../../../services/Terminal.service';
import { Router } from '@angular/router';

describe('CreateTerminalComponent', () => {
  let component: CreateTerminalComponent;
  let fixture: ComponentFixture<CreateTerminalComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateTerminalComponent
      ],
      providers: [
        TerminalService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateTerminalComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});