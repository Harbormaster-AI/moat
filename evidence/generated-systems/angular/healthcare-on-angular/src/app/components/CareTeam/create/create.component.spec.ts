
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateCareTeamComponent } from './create.component';
import { CareTeamService } from '../../../services/CareTeam.service';
import { Router } from '@angular/router';

describe('CreateCareTeamComponent', () => {
  let component: CreateCareTeamComponent;
  let fixture: ComponentFixture<CreateCareTeamComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateCareTeamComponent
      ],
      providers: [
        CareTeamService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateCareTeamComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});