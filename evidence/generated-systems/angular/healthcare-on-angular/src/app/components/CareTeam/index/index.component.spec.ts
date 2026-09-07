
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexCareTeamComponent } from './index.component';
import { CareTeamService } from '../../../services/CareTeam.service';

describe('IndexCareTeamComponent', () => {
  let component: IndexCareTeamComponent;
  let fixture: ComponentFixture<IndexCareTeamComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexCareTeamComponent
      ],
      providers: [
        CareTeamService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexCareTeamComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});