
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexCampaignComponent } from './index.component';
import { CampaignService } from '../../../services/Campaign.service';

describe('IndexCampaignComponent', () => {
  let component: IndexCampaignComponent;
  let fixture: ComponentFixture<IndexCampaignComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexCampaignComponent
      ],
      providers: [
        CampaignService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexCampaignComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});